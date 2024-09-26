//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

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
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Middlewares) (*kratos.App, func(), error) {
	panic(wire.Build(middlewares.ProviderSet, router.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
