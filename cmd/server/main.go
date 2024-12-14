package main

import (
	"context"
	"log"
	"net"

	"github.com/tsunakit99/ankylo-cup-backend/internal/api/handlers"
	"github.com/tsunakit99/ankylo-cup-backend/internal/auth"
	"github.com/tsunakit99/ankylo-cup-backend/internal/config"
	pbg "github.com/tsunakit99/ankylo-cup-backend/internal/pb/game"
	pbr "github.com/tsunakit99/ankylo-cup-backend/internal/pb/room"
	pbs "github.com/tsunakit99/ankylo-cup-backend/internal/pb/score"
	pbu "github.com/tsunakit99/ankylo-cup-backend/internal/pb/user"
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
		grpc.UnaryInterceptor(handlers.AuthInterceptor(authClient)),
	)
	reflection.Register(s)

	userService := handlers.NewUserServiceServer(authClient, dbConn)
	pbu.RegisterUserServiceServer(s, userService)
	roomService := handlers.NewRoomServiceServer(dbConn)
	pbr.RegisterRoomServiceServer(s, roomService)
	gameService := handlers.NewGameServiceServer(dbConn)
	pbg.RegisterGameServiceServer(s, gameService)
	scoreService := handlers.NewScoreServiceServer(dbConn)
	pbs.RegisterScoreServiceServer(s, scoreService)


	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
