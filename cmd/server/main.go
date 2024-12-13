package main

import (
	"log"
	"net"

	grpcapi "github.com/tsunakit99/ankylo-cup-backend/internal/api/grpc"
	"github.com/tsunakit99/ankylo-cup-backend/internal/config"
	"github.com/tsunakit99/ankylo-cup-backend/internal/db"
	"github.com/tsunakit99/ankylo-cup-backend/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// データベースに接続
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	log.Println("DB connected successfully")
	defer database.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userService := &grpcapi.UserServiceServer{}
	pb.RegisterUserServiceServer(s, userService)

	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
