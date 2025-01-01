package auth_adapter

import (
	"context"
	"github.com/ZetNetwork/Protos/pkg/auth_v1"
	"github.com/ZetNetwork/Users/internal/domain/ports"
	"github.com/ZetNetwork/Users/pkg/logger"
)

type authClient struct {
	client auth_v1.AuthV1Client
}

func NewAuthClient(
	client auth_v1.AuthV1Client,
) ports.IAuthClient {
	return &authClient{
		client: client,
	}
}

func (a authClient) ValidateToken(ctx context.Context, accessToken string) error {
	_, err := a.client.ValidateToken(ctx, &auth_v1.ValidateTokenRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		logger.LoggerFromContext(ctx).
			Error("failed to validate token: " + err.Error())
		return err
	}
	return nil
}
