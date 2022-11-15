package di

import (
	"fmt"
	"git.epam.com/ryan_wang/go-web-service/config"
	"git.epam.com/ryan_wang/go-web-service/internal/controllers"
	"git.epam.com/ryan_wang/go-web-service/internal/repositories"
	"git.epam.com/ryan_wang/go-web-service/internal/services"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//func InitWebServerContext() *webserver.Context {
//	wire.Build(
//		config.LoadConfig,
//		repositories.NewRecordRepository,
//		services.NewRecordService,
//		controllers.NewRecordController,
//		webserver.NewServerContext)
//	return &webserver.Context{}
//}

func InitRecordController() *controllers.RecordController {
	wire.Build(
		config.LoadConfig,
		NewDBConnection,
		repositories.NewRecordRepository,
		services.NewRecordService,
		controllers.NewRecordController)
	return &controllers.RecordController{}
}

func NewDBConnection(cfg *config.Configuration) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.UserName,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return db
}
