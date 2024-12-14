package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/tsunakit99/ankylo-cup-backend/internal/models"
	pb "github.com/tsunakit99/ankylo-cup-backend/internal/pb/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RoomServiceServer struct {
	pb.UnimplementedRoomServiceServer
	DB *gorm.DB
}

func NewRoomServiceServer(db *gorm.DB) *RoomServiceServer {
	return &RoomServiceServer{DB: db}
}

func (s *RoomServiceServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	uidValue := ctx.Value("user_uid")
	if uidValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "no user uid found in context")
	}
	uid := uidValue.(string)

	room := models.Room{
		GameID:    int(req.GameId),
		Player1ID: uid,
		Status:    "waiting",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.DB.Create(&room).Error; err != nil {
		log.Printf("Failed to create room: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create room")
	}

	return &pb.CreateRoomResponse{
		RoomId: int32(room.ID),
		Status: "waiting",
	}, nil
}

func (s *RoomServiceServer) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	uidValue := ctx.Value("user_uid")
	if uidValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "no user uid found")
	}
	uid := uidValue.(string)

	var room models.Room
	if err := s.DB.First(&room, "id = ?", req.RoomId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "room not found")
		}
		return nil, status.Errorf(codes.Internal, "db error")
	}

	if room.Status != "waiting" {
		return nil, status.Errorf(codes.FailedPrecondition, "room not available")
	}

	room.Player2ID = uid
	room.Status = "playing"
	room.UpdatedAt = time.Now()

	if err := s.DB.Save(&room).Error; err != nil {
		log.Printf("Failed to update room: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to join room")
	}

	return &pb.JoinRoomResponse{
		RoomId: int32(room.ID),
		Status: "playing",
	}, nil
}

// 双方向ストリーミング例
func (s *RoomServiceServer) StartGame(stream pb.RoomService_StartGameServer) error {
	// 双方向ストリーミング: client->server, server->client
	for {
		in, err := stream.Recv()
		if err != nil {
			// クライアントがストリームを閉じた場合など
			return err
		}
		// inには"action", "data"が入っているとする
		// サーバーが処理して返信する場合:
		if err := stream.Send(&pb.GameAction{
			Action: "ack",
			Data:   "Received: " + in.Data,
		}); err != nil {
			return err
		}
	}
}
