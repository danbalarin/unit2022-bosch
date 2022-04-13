//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/unit2022-bosch/teapot/backend/internal/app/orm"
	"github.com/unit2022-bosch/teapot/backend/internal/app/webserver"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
)

func BuildWebServer() (*webserver.App, error) {
	wire.Build(
		orm.NewDB,
		orm.NewDbConfig,
		orm.NewMigrator,
		webserver.NewApp,
		webserver.NewRouter,
		webserver.NewWebConfig,
		auth.NewAuthRestController,
		auth.NewAuthService,
		auth.NewAuthGormRepository,
		auth.NewAuthConfig,
		auth.NewAuthDbSeeder,
	)
	return &webserver.App{}, nil
}
