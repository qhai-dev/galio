package galio

import (
	"context"
	"net/http"
	"sync"

	"google.golang.org/grpc"
)

type Application struct {
	conf   any
	ctx    context.Context
	cancel context.CancelFunc

	gs *grpc.Server
	rs *http.Server

	mu sync.Mutex

	startHooks []func(ctx context.Context) error
	stopHooks  []func(ctx context.Context) error
}

func (app *Application) StartHook(hook func(ctx context.Context) error) {
	app.startHooks = append(app.startHooks, hook)
}

func (app *Application) StopHook(hook func(ctx context.Context) error) {
	app.stopHooks = append(app.stopHooks, hook)
}

func NewApplication() *Application {
	ctx, cancel := context.WithCancel(context.Background())

	return &Application{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (app *Application) runStartHooks() error {
	for _, hook := range app.startHooks {
		if err := hook(app.ctx); err != nil {
			return err
		}
	}
	return nil
}
