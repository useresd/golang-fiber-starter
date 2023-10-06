package services

import (
	"context"

	"github.com/useresd/golang-fiber-starter/internal/models"
	"github.com/useresd/golang-fiber-starter/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) FindMany(ctx context.Context) ([]*models.User, error) {
	return s.userRepository.FindMany(ctx)
}

func (s *UserService) Store(ctx context.Context, user *models.User) error {
	return s.userRepository.Store(ctx, user)
}
