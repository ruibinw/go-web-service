package webserver

import (
	"fmt"
	"git.epam.com/ryan_wang/go-web-service/config"
	"git.epam.com/ryan_wang/go-web-service/internal/controllers"
	"git.epam.com/ryan_wang/go-web-service/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Context struct {
	Router *echo.Echo
	Config *config.Configuration
}

func NewServerContext(cfg *config.Configuration) *Context {
	return &Context{
		Router: initRouter(),
		Config: cfg,
	}
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

func initRouter() *echo.Echo {
	router := echo.New()
	router.Validator = utils.NewRequestValidator(validator.New())
	router.Use(middleware.Logger()) //to log request info
	router.Use(middleware.CORS())   //to allow cross-domain access
	return router
}

func (ctx *Context) RegisterRecordApiPaths(ctrl controllers.RecordController) {
	ctx.Router.POST("/records", ctrl.Create)
	ctx.Router.PUT("/records/:id", ctrl.Update)
	ctx.Router.DELETE("/records/:id", ctrl.Delete)
	ctx.Router.GET("/records/:id", ctrl.Get)
	ctx.Router.GET("/records", ctrl.Query)
}

func (ctx *Context) Start() {
	ctx.Router.Logger.Fatal(
		ctx.Router.Start(":" + ctx.Config.Server.Port))
}
