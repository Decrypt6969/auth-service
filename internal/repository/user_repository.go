package repository

import (
	"context"
	"time"

	"github.com/decrypt6969/auth-service/internal/db"
	"github.com/decrypt6969/auth-service/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.DB.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepo) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var user model.User
	err := db.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
