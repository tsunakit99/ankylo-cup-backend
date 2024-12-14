package handlers

import (
	"context"

	"github.com/tsunakit99/ankylo-cup-backend/internal/models"
	pb "github.com/tsunakit99/ankylo-cup-backend/internal/pb/game"
	pb2 "github.com/tsunakit99/ankylo-cup-backend/internal/pb/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GameServiceServer struct {
	pb.UnimplementedGameServiceServer
	DB *gorm.DB
}

func NewGameServiceServer(db *gorm.DB) *GameServiceServer {
	return &GameServiceServer{DB: db}
}

func (s *GameServiceServer) GetGameList(ctx context.Context, req *pb.GetGameListRequest) (*pb.GetGameListResponse, error) {
	var games []models.Game
	if err := s.DB.Find(&games).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "db error")
	}

	return &pb.GetGameListResponse{
		Games: convertGamesToPb(games),
	}, nil
}

func (s *GameServiceServer) GetGameById(ctx context.Context, req *pb.GetGameByIdRequest) (*pb.GetGameByIdResponse, error) {
	var g models.Game
	if err := s.DB.First(&g, "id = ?", req.GameId).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "game not found")
	}
	return &pb.GetGameByIdResponse{
		Game: convertGameToPb(g),
	}, nil
}

func convertGamesToPb(gs []models.Game) []*pb2.Game {
	var pbGames []*pb2.Game
	for _, g := range gs {
		pbGames = append(pbGames, convertGameToPb(g))
	}
	return pbGames
}

func convertGameToPb(g models.Game) *pb2.Game {
	return &pb2.Game{
		Id:         int32(g.ID),
		Name:       g.Name,
		MaxPlayers: int32(g.MaxPlayers),
		// created_at, updated_at省略
	}
}
