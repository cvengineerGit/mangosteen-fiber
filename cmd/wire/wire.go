//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/server"
	"github.com/kalougata/bookkeeping/internal/service"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		data.NewData,
		service.NewUserService,
		service.NewTagService,
		controller.NewAuthController,
		controller.NewTagController,
		server.NewHTTPServer,
		server.NewServer,
	))
}
