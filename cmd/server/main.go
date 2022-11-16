package main

import (
	"git.epam.com/ryan_wang/go-web-service/config"
	_ "git.epam.com/ryan_wang/go-web-service/docs"
	"git.epam.com/ryan_wang/go-web-service/internal/controllers"
	"git.epam.com/ryan_wang/go-web-service/internal/repositories"
	"git.epam.com/ryan_wang/go-web-service/internal/services"
	"git.epam.com/ryan_wang/go-web-service/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title 			Go Web Service
// @version 		1.0
// @description 	A simple REST Web service written in Go that supports CRUD operations.
// @contact.name 	Ryan_Wang
// @contact.email 	Ryan_Wang@epam.com
// @host 			localhost:8080
// @BasePath 		/
func main() {
	cfg := config.GetConfig()
	db := cfg.OpenDBConnection()
	recordRepo := repositories.NewRecordRepository(db)
	recordSrv := services.NewRecordService(recordRepo)
	recordCtrl := controllers.NewRecordController(recordSrv)

	server := echo.New()
	server.Validator = utils.NewRequestValidator(validator.New())
	server.Use(middleware.Logger()) //to log request info
	server.Use(middleware.CORS())   //to allow cross-domain access

	server.GET("/", index)
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.POST("/records", recordCtrl.Create)
	server.PUT("/records/:id", recordCtrl.Update)
	server.DELETE("/records/:id", recordCtrl.Delete)
	server.GET("/records/:id", recordCtrl.Get)
	server.GET("/records", recordCtrl.Query)

	server.Logger.Fatal(server.Start(":" + cfg.Server.Port))
}

func index(c echo.Context) error {
	return c.HTML(http.StatusOK,
		`Go Web Service is running. <a href="swagger/index.html">[See API definition]</a>`,
	)
}
