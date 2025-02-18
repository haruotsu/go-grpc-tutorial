package server

import (
	"context"
	"errors"

	"github.com/haruotsu/grpc-test/internal/pb"
	"github.com/haruotsu/grpc-test/internal/service"
)

// UserServer は pb.UserServiceServer インターフェースの実装です。
type UserServer struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

// NewUserServer は新しい UserServer を生成します。
func NewUserServer(us service.UserService) *UserServer {
	return &UserServer{userService: us}
}

// GetUser は gRPC の GetUser エンドポイントを実装します。
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
