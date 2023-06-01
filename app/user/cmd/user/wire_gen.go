// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/conf"
	"beetle/app/user/internal/data"
	"beetle/app/user/internal/server"
	"beetle/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(auth, userRepo, logger)
	friendRepo := data.NewFriendRepo(dataData, logger)
	friendUseCase := biz.NewFriendUseCase(auth, friendRepo, logger)
	userService := service.NewUserService(auth, userUseCase, friendUseCase)
	grpcServer := server.NewGRPCServer(confServer, auth, userService, logger)
	registrar := data.NewRegistrar()
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
