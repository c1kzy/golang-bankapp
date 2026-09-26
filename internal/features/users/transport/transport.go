package core_user_transport

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	PatchUser(ctx context.Context, patch domain.UserPatch, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}
