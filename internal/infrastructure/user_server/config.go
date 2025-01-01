package user_server

import (
	"fmt"
	"os"
)

const (
	_ServerHost = "USER_SERVER_HOST"
	_ServerPort = "USER_SERVER_PORT"
)

type IUserServerConfig interface {
	Address() string
}

type UserServerConfig struct {
	host string
	port string
}

func NewUserServerConfig() (IUserServerConfig, error) {
	host := os.Getenv(_ServerHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("env %s is empty", _ServerHost)
	}

	port := os.Getenv(_ServerPort)
	if len(port) == 0 {
		return nil, fmt.Errorf("env %s is empty", _ServerPort)
	}

	return &UserServerConfig{
		host: host,
		port: port,
	}, nil
}

func (u UserServerConfig) Address() string {
	return u.host + ":" + u.port
}
