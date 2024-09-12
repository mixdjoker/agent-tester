package app

import (
	"context"

	"github.com/mixdjoker/agent-tester/internal/api"
	"github.com/mixdjoker/agent-tester/internal/config"
	"github.com/mixdjoker/agent-tester/internal/logger"
	"github.com/mixdjoker/agent-tester/internal/service"
	"github.com/mixdjoker/agent-tester/internal/service/task"
)

type serviceProvider struct {
	certs     config.Certificator
	apiConfig config.APIConfiger

	taskService service.TesterService

	apiMux *api.Mux
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) APIConfig() config.APIConfiger {
	if s.apiConfig == nil {
		cfg, err := config.NewApiConfig()
		if err != nil {
			logger.Fatalf(context.Background(), err, "failed to get API Config")
		}

		s.apiConfig = cfg
	}

	return s.apiConfig
}

func (s *serviceProvider) APIMux() *api.Mux {
	if s.apiMux == nil {
		mux := api.NewMux(s.TesterService())
		s.apiMux = mux
	}

	return s.apiMux
}

func (s *serviceProvider) TesterService() service.TesterService {
	if s.taskService == nil {
		ts := task.NewService()
		s.taskService = ts
	}

	return s.taskService
}

func (s *serviceProvider) Certs() config.Certificator {
	if s.certs == nil {
		certs, err := config.NewCertConfig()
		if err != nil {
			logger.Fatalf(context.Background(), err, "failed to load certs files")
		}

		s.certs = certs
	}

	return s.certs
}
