package ports

import "context"

type IAuthClient interface {
	ValidateToken(ctx context.Context, accessToken string) error
}
