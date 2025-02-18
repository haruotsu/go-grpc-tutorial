package service

import (
	"context"

	"github.com/haruotsu/grpc-test/internal/model"
	"github.com/haruotsu/grpc-test/internal/repository"
)

// UserServiceは、ユーザーに関するビジネスロジックを提供するインターフェース。
type UserService interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

// userServiceはUserServiceインターフェースの実装です。
type userService struct {
	repo repository.UserRepository
}

// NewUserServiceは新しいUserService を生成します。
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// GetUserは、repository層を利用してユーザー情報を取得します。
func (s *userService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
