package server

import (
	"github.com/qhai-dev/galio/rpc/account/server/handler"
)

func NewRPCServer(handler *handler.AccountService) error {
	// srv := rpc.NewServer()

	// handler.Handler()

	// return srv
	handler.Handler()

	return nil
}
