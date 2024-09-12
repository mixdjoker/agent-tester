package app

import (
	"context"

	"github.com/mixdjoker/agent-tester/internal/closer"
	"github.com/mixdjoker/agent-tester/internal/config"
	"github.com/mixdjoker/agent-tester/internal/logger"
	"github.com/quic-go/quic-go/http3"
	"github.com/spf13/pflag"
)

// App is a struct that holds all the application dependencies
type App struct {
	sp *serviceProvider
	// Servers
	api *http3.Server
}

// NewApp is a function that returns a new instance of the App struct
func NewApp(ctx context.Context) (*App, error) {
	a := App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		// Init functions
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initApiServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envFilePath := pflag.StringP("env", "e", ".env", ".env file path")
	pflag.Parse()

	if err := config.Load(*envFilePath); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	cfg, err := config.NewLoggerConfig()
	if err != nil {
		return err
	}

	logger.SetupLoggerLevel(cfg.Level(), cfg.Format())
	logger.Info(ctx).Msg("Logger initialized")

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider()
	return nil
}

func (a *App) initApiServer(_ context.Context) error {
	a.api = &http3.Server{
		Addr:    a.sp.APIConfig().Address(),
		Handler: a.sp.APIMux(),
	}

	return nil
}

// Run starts initialised application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx := context.Background()

	if err := a.api.ListenAndServeTLS(a.sp.Certs().Cert(), a.sp.Certs().Key()); err != nil {
		logger.Fatalf(ctx, err, "http/3 server failes")
	}

	return nil
}
