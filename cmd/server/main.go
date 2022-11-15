package main

import (
	"fmt"
	"git.epam.com/ryan_wang/crud-demo/config"
	"git.epam.com/ryan_wang/crud-demo/internal/controllers"
	"git.epam.com/ryan_wang/crud-demo/internal/repositories"
	"git.epam.com/ryan_wang/crud-demo/internal/services"
	"git.epam.com/ryan_wang/crud-demo/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

// @title CRUD Demo in Go
// @version 1.0
// @description A simple REST Web service that supports CRUD operations.

// @contact.name Ryan_Wang
// @contact.email Ryan_Wang@epam.com

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.GetConfig()
	db := NewDBConnection(cfg)
	recordRepo := repositories.NewRecordRepository(db)
	recordSrv := services.NewRecordService(recordRepo)
	recordCtrl := controllers.NewRecordController(recordSrv)

	server := echo.New()
	server.Validator = utils.NewRequestValidator(validator.New())
	server.Use(middleware.Logger()) //to log request info

	server.GET("/", index)
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.POST("/records", recordCtrl.Create)
	server.PUT("/records/:id", recordCtrl.Update)
	server.DELETE("/records/:id", recordCtrl.Delete)
	server.GET("/records/:id", recordCtrl.Get)
	server.GET("/records", recordCtrl.Query)

	server.Logger.Fatal(server.Start(":" + cfg.Server.Port))
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

func index(c echo.Context) error {
	return c.HTML(http.StatusOK,
		`CRUD Demo web service is running. <a href="swagger/index.html">[See API definition]</a>`,
	)
}
