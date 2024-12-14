package handlers

import (
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define custom context keys to avoid collisions
type contextKey string

const (
	userUIDKey contextKey = "user_uid"
	roomIDKey  contextKey = "room_id"
)

// AuthInterceptor はUnary RPC用のInterceptorです。
func AuthInterceptor(authClient *auth.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "no metadata found")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		parts := strings.SplitN(tokens[0], " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization format")
		}

		idToken := parts[1]
		token, verr := authClient.VerifyIDToken(ctx, idToken)
		if verr != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid ID token: %v", verr)
		}

		// Save user UID in context with a custom key
		ctx = context.WithValue(ctx, userUIDKey, token.UID)
		return handler(ctx, req)
	}
}

// StreamAuthInterceptor はストリーミングRPC用のInterceptorです。
func StreamAuthInterceptor(authClient *auth.Client) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		// Authorizationヘッダーの取得
		var idToken string
		if authHeaders, exists := md["authorization"]; exists && len(authHeaders) > 0 {
			parts := strings.SplitN(authHeaders[0], " ", 2)
			if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
				idToken = parts[1]
			}
		}

		if idToken == "" {
			return status.Errorf(codes.Unauthenticated, "missing authorization token")
		}

		// IDトークンの検証
		token, err := authClient.VerifyIDToken(ss.Context(), idToken)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid ID token: %v", err)
		}

		// UIDをコンテキストに保存
		ctx := context.WithValue(ss.Context(), userUIDKey, token.UID)

		// ルームIDの取得
		var roomID int32
		if roomIDs, exists := md["room-id"]; exists && len(roomIDs) > 0 {
			_, err := fmt.Sscanf(roomIDs[0], "%d", &roomID)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "invalid room ID")
			}
			ctx = context.WithValue(ctx, roomIDKey, roomID)
		} else {
			// ルームIDが必要なRPCの場合はエラーを返す
			return status.Errorf(codes.InvalidArgument, "missing room ID")
		}

		// 新しいストリームコンテキストを作成
		wrapped := NewContextStream(ctx, ss)

		// ハンドラを呼び出す
		return handler(srv, wrapped)
	}
}

// contextStream は新しいコンテキストを持つ ServerStream を実装します。
type contextStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *contextStream) Context() context.Context {
	return s.ctx
}

// NewContextStream は新しいコンテキストを持つ ServerStream を作成します。
func NewContextStream(ctx context.Context, ss grpc.ServerStream) grpc.ServerStream {
	return &contextStream{ServerStream: ss, ctx: ctx}
}
