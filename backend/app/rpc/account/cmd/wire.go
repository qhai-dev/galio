//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/qhai-dev/galio/library/galio"
	"github.com/qhai-dev/galio/rpc/account/server"
)

func initializer(*galio.Application) error {
	wire.Build(
		server.ProviderSet,
	)

	return nil
}
