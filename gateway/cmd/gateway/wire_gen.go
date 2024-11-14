// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gateway/internal/biz"
	"gateway/internal/conf"
	"gateway/internal/data"
	"gateway/internal/middlewares"
	"gateway/internal/router"
	"gateway/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, confMiddlewares *conf.Middlewares) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	gatewayBlogRepo := data.NewGatewayBlogRepo(dataData, logger)
	gatewayUploadRepo := data.NewGatewayUploadRepo(dataData, logger)
	gatewayUserRepo := data.NewGatewayUserRepo(dataData, logger)
	gatewayBlogUsecase := biz.NewGatewayBlogUsecase(gatewayBlogRepo, gatewayUploadRepo, gatewayUserRepo)
	gatewayCategoryRepo := data.NewGatewayCategoryRepo(dataData, logger)
	gatewayCategoryUsecase := biz.NewGatewayCategoryUsecase(gatewayCategoryRepo)
	gatewayUserUsecase := biz.NewGatewayUserUsecase(gatewayUserRepo)
	gatewayUploadUsecase := biz.NewGatewayUploadUsecase(gatewayUploadRepo)
	gatewayStatUserRepo := data.NewGatewayStatUserRepo(dataData, logger)
	gatewayStatUserUsecase := biz.NewGatewayStatUserUsecase(gatewayStatUserRepo)
	mids := middlewares.NewMids(confMiddlewares, logger)
	engine := router.NewGinRouter(gatewayBlogUsecase, gatewayCategoryUsecase, gatewayUserUsecase, gatewayUploadUsecase, gatewayStatUserUsecase, mids)
	httpServer := server.NewHTTPServer(confServer, logger, engine)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
