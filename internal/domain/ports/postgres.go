package ports

import (
	"context"

	"github.com/ZetNetwork/Users/internal/domain/models/entities"
)

type IUserRepository interface {
	SetUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, email string, user entities.User) error
	GetUser(ctx context.Context, email string) (*entities.User, error)
}
