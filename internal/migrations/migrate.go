package migrations

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"

	"github.com/ZetNetwork/Users/internal/infrastructure/database/go_postgres"
)

var (
	//go:embed migrations/users/*.sql
	migrations embed.FS
)

func MigrateDB(db *go_postgres.PostgreClient) error {
	if err := migrate(db, "migrations/users"); err != nil {
		return fmt.Errorf("migrate: %v", err)
	}
	return nil
}

func migrate(db *go_postgres.PostgreClient, dir string) error {
	goose.SetBaseFS(migrations)
	if err := goose.Up(db.SqlDB().DB, dir); err != nil {
		return fmt.Errorf("goose up: %v", err)
	}
	return nil
}
