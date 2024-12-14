package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	"firebase.google.com/go/auth"
	"github.com/tsunakit99/ankylo-cup-backend/internal/models"
	pb "github.com/tsunakit99/ankylo-cup-backend/internal/pb/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	AuthClient *auth.Client
	DB         *gorm.DB
}

func NewUserServiceServer(authClient *auth.Client, db *gorm.DB) *UserServiceServer {
	return &UserServiceServer{
		AuthClient: authClient,
		DB:         db,
	}
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// reqは空メッセージ
	uidValue := ctx.Value("user_uid")
	if uidValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "no user uid found in context")
	}
	uid := uidValue.(string)

	u, err := s.AuthClient.GetUser(ctx, uid)
	if err != nil {
		log.Printf("Failed to get user from firebase: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user info from firebase")
	}

	var user models.User
	err = s.DB.First(&user, "id = ?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// display_name, avatar_urlはFirebaseに接続して取得するのが要件。
		// ここでは Firebase User RecordからdisplayName, photoURLを使用可能。
		displayName := u.DisplayName
		avatarURL := u.PhotoURL

		user = models.User{
			ID:          uid,
			DisplayName: displayName,
			AvatarURL:   avatarURL,
			Coin: 0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err2 := s.DB.Create(&user).Error; err2 != nil {
			log.Printf("DB create user error: %v", err2)
			return nil, status.Errorf(codes.Internal, "failed to create user in DB")
		}
	} else if err != nil {
		log.Printf("DB error: %v", err)
		return nil, status.Errorf(codes.Internal, "db error")
	}

	return &pb.SignUpResponse{
		UserId:      user.ID,
		DisplayName: user.DisplayName,
		AvatarUrl:   user.AvatarURL,
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user models.User
	if err := s.DB.First(&user, "id = ?", req.GetUserId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "db error")
	}

	return &pb.GetUserResponse{
		UserId:      user.ID,
		DisplayName: user.DisplayName,
	}, nil
}

// AddCoin API
func (s *UserServiceServer) AddCoin(ctx context.Context, req *pb.AddCoinRequest) (*pb.AddCoinResponse, error) {
	uidValue := ctx.Value("user_uid")
	if uidValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "no user uid found")
	}
	uid := uidValue.(string)

	var user models.User
	if err := s.DB.First(&user, "id = ?", uid).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	user.Coin += req.Amount
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add coin")
	}

	return &pb.AddCoinResponse{
		Message: "Coin added",
	}, nil
}

// GetTop10ByCoin API
func (s *UserServiceServer) GetTop10ByCoin(ctx context.Context, req *pb.GetTop10ByCoinRequest) (*pb.GetTop10ByCoinResponse, error) {
	var users []models.User
	if err := s.DB.Order("coin DESC NULLS LAST").Limit(10).Find(&users).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "db error")
	}

	var topUsers []*pb.TopUserByCoin
	for _, u := range users {
		topUsers = append(topUsers, &pb.TopUserByCoin{
			UserId:      u.ID,
			DisplayName: u.DisplayName,
			Coin:        int32(u.Coin), // Coinがnil不可なら0デフォルト
		})
	}

	return &pb.GetTop10ByCoinResponse{Users: topUsers}, nil
}

