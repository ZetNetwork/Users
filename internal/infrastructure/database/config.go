package database

import (
	"fmt"
	"os"
)

const (
	_DSN    = "PG_DSN"
	_DRIVER = "PG_DRIVER"
)

type IPGConfig interface {
	GetDSN() string
	GetDriver() string
}

type PGConfig struct {
	dsn    string
	driver string
}

func NewPGConfig() (IPGConfig, error) {
	dsn := os.Getenv(_DSN)
	driver := os.Getenv(_DRIVER)

	if len(dsn) == 0 || len(driver) == 0 {
		return nil, fmt.Errorf("env %s or %s is empty", _DSN, _DRIVER)
	}

	return &PGConfig{
		dsn:    dsn,
		driver: driver,
	}, nil
}

func (c *PGConfig) GetDSN() string {
	return c.dsn
}

func (c *PGConfig) GetDriver() string {
	return c.driver
}
