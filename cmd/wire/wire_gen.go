// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/server"
	"github.com/kalougata/bookkeeping/internal/service"
)

// Injectors from wire.go:

func NewApp() (*server.Server, func(), error) {
	dataData, cleanup, err := data.NewData()
	if err != nil {
		return nil, nil, err
	}
	userService := service.NewUserService(dataData)
	authController := controller.NewAuthController(userService)
	app := server.NewHTTPServer(authController)
	serverServer := server.NewServer(app)
	return serverServer, func() {
		cleanup()
	}, nil
}
