package handlers

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

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

		ctx = context.WithValue(ctx, "user_uid", token.UID)
		return handler(ctx, req)
	}
}
