package config

import (
	"fmt"
	"net"
	"os"
)

const (
	apiHostEnv = "TESTER_API_HOST"
	apiPortEnv = "TESTER_API_PORT"
)

type APIConfiger interface {
	Address() string
}

type apiConfig struct {
	host string
	port string
}

func NewApiConfig() (APIConfiger, error) {
	cfg := apiConfig{
		host: os.Getenv(apiHostEnv),
		port: os.Getenv(apiPortEnv),
	}

	if cfg.port == "" {
		return nil, fmt.Errorf("can't get %s", apiPortEnv)
	}

	return &cfg, nil
}

func (a *apiConfig) Address() string {
	return net.JoinHostPort(a.host, a.port)
}
