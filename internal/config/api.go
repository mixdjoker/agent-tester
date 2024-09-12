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

// APIConfiger ...
type APIConfiger interface {
	Address() string
}

type apiConfig struct {
	host string
	port string
}

// NewAPIConfig ...
func NewAPIConfig() (APIConfiger, error) {
	cfg := apiConfig{
		host: os.Getenv(apiHostEnv),
		port: os.Getenv(apiPortEnv),
	}

	if cfg.port == "" {
		return nil, fmt.Errorf("can't get %s", apiPortEnv)
	}

	return &cfg, nil
}

// Address ...
func (a *apiConfig) Address() string {
	return net.JoinHostPort(a.host, a.port)
}
