package app

import "context"

// App is a struct that holds all the application dependencies
type App struct {
	// Service Provider
	// Servers
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
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}
