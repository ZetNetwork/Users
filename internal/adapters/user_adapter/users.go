package user_adapter

import (
	"context"
	"fmt"
	"github.com/ZetNetwork/Protos/pkg/users_v1"
	"github.com/ZetNetwork/Users/internal/domain/models/dto"
	"github.com/ZetNetwork/Users/internal/domain/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersServer struct {
	users_v1.UnimplementedUsersV1Server
	uc usecase.IUserUseCase
}

func NewUserServer(uc usecase.IUserUseCase) *UsersServer {
	return &UsersServer{
		uc: uc,
	}
}

func (s *UsersServer) CreateUser(ctx context.Context, req *users_v1.CreateUserRequest) (*emptypb.Empty, error) {
	token := req.AccessToken
	result, err := s.uc.ValidateToken(ctx, token)

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	if !result {
		return nil, fmt.Errorf("invalid token")
	}

	user := dto.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
	}

	err = s.uc.SetUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to set user: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *UsersServer) UpdateUser(ctx context.Context, req *users_v1.UpdateUserRequest) (*emptypb.Empty, error) {
	token := req.AccessToken
	result, err := s.uc.ValidateToken(ctx, token)

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	if !result {
		return nil, fmt.Errorf("invalid token")
	}

	user := dto.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
	}

	err = s.uc.UpdateUser(ctx, req.Email, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *UsersServer) GetUser(ctx context.Context, req *users_v1.GetUserRequest) (*users_v1.GetUserResponse, error) {
	user, err := s.uc.GetUser(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &users_v1.GetUserResponse{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, req *users_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	token := req.AccessToken
	result, err := s.uc.ValidateToken(ctx, token)

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	if !result {
		return nil, fmt.Errorf("invalid token")
	}

	err = s.uc.DeleteUser(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}
	return &emptypb.Empty{}, nil
}
