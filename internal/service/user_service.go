package service

import (
	"errors"

	"github.com/decrypt6969/auth-service/internal/model"
	"github.com/decrypt6969/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *model.User) error {
	existing, _ := s.repo.GetByEmail(user.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return s.repo.Create(user)
}
