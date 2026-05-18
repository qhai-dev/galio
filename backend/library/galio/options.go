package galio

import "context"

type Option func(*options)

type options struct {
	startHook []func(ctx context.Context) error
	stopHook  []func(ctx context.Context) error
}

func WithStartHook(hook func(ctx context.Context) error) Option {
	return func(o *options) {
		o.startHook = append(o.startHook, hook)
	}
}

func WithStopHook(hook func(ctx context.Context) error) Option {
	return func(o *options) {
		o.stopHook = append(o.stopHook, hook)
	}
}
