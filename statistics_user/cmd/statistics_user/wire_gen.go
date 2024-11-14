// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"statistics_user/internal/biz"
	"statistics_user/internal/conf"
	"statistics_user/internal/data"
	"statistics_user/internal/server"
	"statistics_user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	statUserRepo := data.NewStatUserRepo(dataData, logger)
	statUserUsecase := biz.NewStatUserUsecase(statUserRepo)
	statUserService := service.NewStatUserService(statUserUsecase)
	grpcServer := server.NewGRPCServer(confServer, logger, statUserService)
	app := newApp(logger, grpcServer, dataData)
	return app, func() {
		cleanup()
	}, nil
}
