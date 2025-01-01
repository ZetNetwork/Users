package repository

import (
	"context"
	"fmt"
	"github.com/ZetNetwork/Users/internal/domain/models/entities"
	"github.com/ZetNetwork/Users/internal/domain/ports"
	"github.com/ZetNetwork/Users/internal/infrastructure/database/go_postgres"
)

const (
	_TABLE_USERS = "users"
)

type userRepository struct {
	client *go_postgres.PostgreClient
}

func NewUserRepository(
	client *go_postgres.PostgreClient,
) ports.IUserRepository {
	return &userRepository{
		client: client,
	}
}

func (u userRepository) SetUser(ctx context.Context, user entities.User) error {
	qb := u.client.Builder.Insert(_TABLE_USERS).
		Columns(
			"email",
			"password",
			"name",
			"surname",
		).
		Values(
			user.Email,
			user.Password,
			user.Name,
			user.Surname,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	_, err = u.client.SqlDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

func (u userRepository) DeleteUser(ctx context.Context, email string) error {
	qb := u.client.Builder.Delete(_TABLE_USERS).Where("email = ?", email)

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	_, err = u.client.SqlDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

func (u userRepository) UpdateUser(ctx context.Context, email string, user entities.User) error {
	qb := u.client.Builder.Update(_TABLE_USERS).
		Set("name", user.Name).
		Set("surname", user.Surname).
		Where("email = ?", email)

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := u.client.SqlDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %w", err)
	}

	return nil
}

func (u userRepository) GetUser(ctx context.Context, email string) (*entities.User, error) {
	qb := u.client.Builder.Select("*").From(_TABLE_USERS).Where("email = ?", email)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var user *entities.User

	if err := u.client.SqlDB().SelectContext(ctx, &user, query, args...); err != nil {
		return nil, fmt.Errorf("select context: %w", err)
	}

	return user, nil
}
