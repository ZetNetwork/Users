package usecase

import (
	"context"
	"fmt"
	"github.com/ZetNetwork/Users/internal/converters"
	"github.com/ZetNetwork/Users/internal/domain/models/dto"
	"github.com/ZetNetwork/Users/internal/domain/ports"
)

type UserUseCase struct {
	repo       ports.IUserRepository
	authHelper ports.IAuthClient
	converter  converters.UsersConverter
}

func NewUserUseCase(repo ports.IUserRepository, authHelper ports.IAuthClient) *UserUseCase {
	return &UserUseCase{
		repo:       repo,
		authHelper: authHelper,
	}
}

func (u *UserUseCase) SetUser(ctx context.Context, user dto.User) error {
	userEntity := u.converter.UserToEntity(user)
	err := u.repo.SetUser(ctx, userEntity)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	return nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, email string) error {
	err := u.repo.DeleteUser(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, email string, user dto.User) error {
	userEntity := u.converter.UserToEntity(user)
	err := u.repo.UpdateUser(ctx, email, userEntity)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (u *UserUseCase) GetUser(ctx context.Context, email string) (*dto.User, error) {
	userEntity, err := u.repo.GetUser(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	user := u.converter.UserToDTO(*userEntity)
	return &user, nil
}

func (u *UserUseCase) ValidateToken(ctx context.Context, token string) (bool, error) {
	err := u.authHelper.ValidateToken(ctx, token)
	if err != nil {
		return false, fmt.Errorf("failed to validate token: %v", err)
	}
	return true, nil
}
