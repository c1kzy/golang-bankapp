package core_http_service

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	PatchUser(ctx context.Context, patch domain.User, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
