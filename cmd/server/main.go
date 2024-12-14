package main

import (
	"context"
	"log"
	"net"

	"github.com/tsunakit99/ankylo-cup-backend/internal/api/handlers"
	"github.com/tsunakit99/ankylo-cup-backend/internal/auth"
	"github.com/tsunakit99/ankylo-cup-backend/internal/config"
	pb_chinchiro "github.com/tsunakit99/ankylo-cup-backend/internal/pb/chinchiro"
	pb_game "github.com/tsunakit99/ankylo-cup-backend/internal/pb/game"
	pb_room "github.com/tsunakit99/ankylo-cup-backend/internal/pb/room"
	pb_score "github.com/tsunakit99/ankylo-cup-backend/internal/pb/score"
	pb_user "github.com/tsunakit99/ankylo-cup-backend/internal/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()
	app := auth.InitFirebaseApp(ctx, cfg.CredentialsFilePath)
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to get Auth client: %v", err)
	}

	// GORMでのDB接続
	dbConn, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	log.Println("DB connected successfully")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		// ユニキャストRPC用のInterceptor
		grpc.UnaryInterceptor(handlers.AuthInterceptor(authClient)),
		// ストリーミングRPC用のInterceptor
		grpc.StreamInterceptor(handlers.StreamAuthInterceptor(authClient)),
	)
	reflection.Register(s)

	 pb_room.RegisterRoomServiceServer(s, handlers.NewRoomServiceServer(dbConn))
    pb_chinchiro.RegisterChinchiroServiceServer(s, handlers.NewChinchiroServiceServer())
    pb_user.RegisterUserServiceServer(s, handlers.NewUserServiceServer(authClient, dbConn))
    pb_game.RegisterGameServiceServer(s, handlers.NewGameServiceServer(dbConn))
    pb_score.RegisterScoreServiceServer(s, handlers.NewScoreServiceServer(dbConn))

	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
