package auth_server

import (
	"fmt"
	"os"
)

const (
	_AuthHost = "GRPC_AUTH_HOST"
	_AuthPort = "GRPC_AUTH_PORT"
)

type IAuthClientConfig interface {
	Address() string
}

type AuthClientConfig struct {
	port string
	host string
}

func NewAuthClientConfig() (IAuthClientConfig, error) {
	host := os.Getenv(_AuthHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("env %s is empty", _AuthHost)
	}

	port := os.Getenv(_AuthPort)
	if len(port) == 0 {
		return nil, fmt.Errorf("env %s is empty", _AuthPort)
	}

	return &AuthClientConfig{
		host: host,
		port: port,
	}, nil
}

func (a AuthClientConfig) Address() string {
	return a.port + ":" + a.host
}
