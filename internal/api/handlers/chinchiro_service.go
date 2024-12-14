package handlers

import (
	"context"
	"errors"
	"log"
	"sync"

	pb "github.com/tsunakit99/ankylo-cup-backend/internal/pb/chinchiro"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ChinchiroServiceServer は ChinchiroService の実装です。
type ChinchiroServiceServer struct {
	pb.UnimplementedChinchiroServiceServer
	// roomStreams 管理用のマップ（Room ID -> ChinchiroRoom）
	roomStreams map[int32]*ChinchiroRoom
	mu          sync.Mutex
}

// ChinchiroRoom は特定のルーム内のプレイヤーのストリームを管理します。
type ChinchiroRoom struct {
	Player1Stream pb.ChinchiroService_PlayChinchiroServer
	Player2Stream pb.ChinchiroService_PlayChinchiroServer
	CurrentTurn    pb.TurnStatus // 現在のターン
	mu             sync.Mutex
}

// NewChinchiroServiceServer は ChinchiroServiceServer のコンストラクタです。
func NewChinchiroServiceServer() *ChinchiroServiceServer {
	return &ChinchiroServiceServer{
		roomStreams: make(map[int32]*ChinchiroRoom),
	}
}

// PlayChinchiro は双方向ストリーミングRPCの実装です。
func (s *ChinchiroServiceServer) PlayChinchiro(stream pb.ChinchiroService_PlayChinchiroServer) error {
    // コンテキストからルームIDとユーザーIDを取得
    roomID, err := getRoomIDFromContext(stream.Context())
    if err != nil {
        return status.Errorf(codes.InvalidArgument, "missing room ID")
    }

    uidValue := stream.Context().Value("user_uid")
    if uidValue == nil {
        return status.Errorf(codes.Unauthenticated, "no user uid found in context")
    }
    userID := uidValue.(string)

    // ルームを取得または作成
    s.mu.Lock()
    room, exists := s.roomStreams[roomID]
    if (!exists) {
        room = &ChinchiroRoom{
            CurrentTurn: pb.TurnStatus_PLAYER1_TURN, // 初期ターンをPlayer1に設定
        }
        s.roomStreams[roomID] = room
    }
    s.mu.Unlock()

    room.mu.Lock()
    defer room.mu.Unlock()

    var playerNum int32
    if room.Player1Stream == nil {
        room.Player1Stream = stream
        playerNum = 1
        log.Printf("Player1 (%s) joined room %d", userID, roomID)
    } else if room.Player2Stream == nil {
        room.Player2Stream = stream
        playerNum = 2
        log.Printf("Player2 (%s) joined room %d", userID, roomID)
    } else {
        return status.Errorf(codes.ResourceExhausted, "room %d already has two players", roomID)
    }

    // 初期ターン情報を送信
    initialMsg := &pb.ChinchiroMessage{
        UserId:     "system",
        DiceRolls:  []int32{},
        Timestamp:  timestamppb.Now(),
        TurnStatus: room.CurrentTurn,
    }
    if playerNum == 1 && room.Player2Stream != nil {
        if err := room.Player2Stream.Send(initialMsg); err != nil {
            log.Printf("Failed to send initial turn status to Player2 in room %d: %v", roomID, err)
        }
    }
    if playerNum == 2 && room.Player1Stream != nil {
        if err := room.Player1Stream.Send(initialMsg); err != nil {
            log.Printf("Failed to send initial turn status to Player1 in room %d: %v", roomID, err)
        }
    }

    // 一回のキャッチボールで終了するため、forループを削除
    msg, err := stream.Recv()
    if err != nil {
        if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
            log.Printf("Player (%s) in room %d disconnected gracefully", userID, roomID)
        } else {
            log.Printf("Player (%s) in room %d disconnected with error: %v", userID, roomID, err)
        }

        // プレイヤーが切断された場合、相手プレイヤーに通知
        systemMsg := &pb.ChinchiroMessage{
            UserId:     "system",
            DiceRolls:  []int32{},
            Timestamp:  timestamppb.Now(),
            TurnStatus: pb.TurnStatus_UNKNOWN, // ゲーム中断を示すステータス
        }
        if room.Player1Stream != nil && room.Player1Stream != stream {
            room.Player1Stream.Send(systemMsg)
        }
        if room.Player2Stream != nil && room.Player2Stream != stream {
            room.Player2Stream.Send(systemMsg)
        }

        // ルームから削除
        s.mu.Lock()
        delete(s.roomStreams, roomID)
        s.mu.Unlock()

        return err
    }

    // 受信したメッセージを相手プレイヤーに転送
    var targetStream pb.ChinchiroService_PlayChinchiroServer
    var nextTurn pb.TurnStatus

    if playerNum == 1 && room.Player2Stream != nil {
        targetStream = room.Player2Stream
        nextTurn = pb.TurnStatus_PLAYER2_TURN
    } else if playerNum == 2 && room.Player1Stream != nil {
        targetStream = room.Player1Stream
        nextTurn = pb.TurnStatus_PLAYER1_TURN
    } else {
        // 相手プレイヤーがまだ接続していない場合、ターンを変更できない
        nextTurn = room.CurrentTurn
    }

    // ターンステータスを更新
    room.CurrentTurn = nextTurn

    // 送信するメッセージにターンステータスを設定
    responseMsg := &pb.ChinchiroMessage{
        UserId:     msg.UserId,
        DiceRolls:  msg.DiceRolls,
        Timestamp:  timestamppb.Now(),
        TurnStatus: room.CurrentTurn,
    }

    if targetStream != nil {
        if err := targetStream.Send(responseMsg); err != nil {
            log.Printf("Failed to send to opponent in room %d: %v", roomID, err)
            return err
        }
    }

    // ルームを削除しない
    return nil
}

