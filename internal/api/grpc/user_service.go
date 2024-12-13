package grpcapi

import (
	"context"

	"github.com/tsunakit99/ankylo-cup-backend/proto/pb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Placeholder
	return &pb.GetUserResponse{
		UserId:       req.UserId,
		DisplayName: "Sample User",
	}, nil
}
