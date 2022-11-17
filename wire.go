// +build wireinject

package main

import (
	"git.epam.com/ryan_wang/go-web-service/internal/config"
	record2 "git.epam.com/ryan_wang/go-web-service/internal/domains/record"
	"git.epam.com/ryan_wang/go-web-service/internal/web"
	"github.com/google/wire"
)

func InitServer() *web.Server {
	wire.Build(config.GetConfig, web.NewServer)
	return &web.Server{}
}

func InitRecordController() *record2.Controller {
	wire.Build(
		config.NewDBConnection,
		record2.NewRepository,
		record2.NewService,
		record2.NewController)
	return &record2.Controller{}
}
