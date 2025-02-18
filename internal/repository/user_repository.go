package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/haruotsu/grpc-test/internal/model"
)

// UserRepository はユーザー情報取得のためのインターフェースです。
type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository は新しいリポジトリインスタンスを生成します。
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUserByID は SQL クエリを利用して、指定されたIDのユーザー情報を取得します。
func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
