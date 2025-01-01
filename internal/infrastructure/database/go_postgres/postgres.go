package go_postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/ZetNetwork/Users/internal/infrastructure/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreClient struct {
	db      *sqlx.DB
	Builder squirrel.StatementBuilderType
}

func NewPostgresClient(config database.IPGConfig) (*PostgreClient, error) {
	db, err := sqlx.Open(config.GetDriver(), config.GetDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreClient{
		db:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (p *PostgreClient) SqlDB() *sqlx.DB {
	return p.db
}
