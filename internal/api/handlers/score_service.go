package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/tsunakit99/ankylo-cup-backend/internal/models"
	pb "github.com/tsunakit99/ankylo-cup-backend/internal/pb/score"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ScoreServiceServer struct {
	pb.UnimplementedScoreServiceServer
	DB *gorm.DB
}

func NewScoreServiceServer(db *gorm.DB) *ScoreServiceServer {
	return &ScoreServiceServer{DB: db}
}

func (s *ScoreServiceServer) RecordScore(ctx context.Context, req *pb.RecordScoreRequest) (*pb.RecordScoreResponse, error) {
	uidValue := ctx.Value("user_uid")
	if uidValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "no user uid found")
	}
	uid := uidValue.(string)

	score := models.Score{
		UserID:   uid,
		GameID:   int(req.GameId),
		Score:    int(req.Score),
		PlayedAt: time.Now(),
	}

	if err := s.DB.Create(&score).Error; err != nil {
		log.Printf("Failed to record score: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to record score")
	}

	return &pb.RecordScoreResponse{
		Message: "Score recorded",
	}, nil
}

func (s *ScoreServiceServer) GetTop10ScoresByGame(ctx context.Context, req *pb.GetTop10ScoresByGameRequest) (*pb.GetTop10ScoresByGameResponse, error) {
	var scores []models.Score
	if err := s.DB.Order("score DESC").Limit(10).Find(&scores, "game_id = ?", req.GameId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.GetTop10ScoresByGameResponse{}, nil // 0件でも空で返す
		}
		return nil, status.Errorf(codes.Internal, "db error")
	}

	var topScores []*pb.TopScoreEntry
	for _, sc := range scores {
		topScores = append(topScores, &pb.TopScoreEntry{
			UserId: sc.UserID,
			Score:  int32(sc.Score),
		})
	}

	return &pb.GetTop10ScoresByGameResponse{
		Scores: topScores,
	}, nil
}
