package usecase

import (
	"context"

	"github.com/ZetNetwork/Users/internal/domain/models/dto"
)

type IUserUseCase interface {
	SetUser(ctx context.Context, user dto.User) error
	DeleteUser(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, email string, user dto.User) error
	GetUser(ctx context.Context, email string) (*dto.User, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
}
