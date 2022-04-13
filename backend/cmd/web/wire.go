//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/unit2022-bosch/teapot/backend/internal/app/orm"
	"github.com/unit2022-bosch/teapot/backend/internal/app/webserver"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"github.com/unit2022-bosch/teapot/backend/internal/services/items"
	"github.com/unit2022-bosch/teapot/backend/internal/services/journeys"
	"github.com/unit2022-bosch/teapot/backend/internal/services/warehouse"
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
		items.NewController,
		items.NewService,
		items.NewRepository,
		items.NewDbSeeder,
		journeys.NewController,
		journeys.NewService,
		journeys.NewRepository,
		journeys.NewDbSeeder,
		warehouse.NewController,
		warehouse.NewService,
		warehouse.NewRepository,
		warehouse.NewDbSeeder,
	)
	return &webserver.App{}, nil
}
