//go:build wireinject
// +build wireinject

package di

import (
	"git.epam.com/ryan_wang/go-web-service/config"
	"git.epam.com/ryan_wang/go-web-service/internal/controllers"
	"git.epam.com/ryan_wang/go-web-service/internal/repositories"
	"git.epam.com/ryan_wang/go-web-service/internal/services"
	"git.epam.com/ryan_wang/go-web-service/internal/web"
	"github.com/google/wire"
)

func InitServer() *web.Server {
	wire.Build(config.GetConfig, web.NewServer)
	return &web.Server{}
}

func InitRecordController() *controllers.RecordController {
	wire.Build(
		config.NewDBConnection,
		repositories.NewRecordRepository,
		services.NewRecordService,
		controllers.NewRecordController)
	return &controllers.RecordController{}
}
