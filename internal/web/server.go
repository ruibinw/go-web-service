package web

import (
	"git.epam.com/ryan_wang/go-web-service/docs"
	"git.epam.com/ryan_wang/go-web-service/internal/config"
	"git.epam.com/ryan_wang/go-web-service/internal/domains/record"
	"git.epam.com/ryan_wang/go-web-service/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type Server struct {
	Router *echo.Echo
	Config *config.Configuration
}

func NewServer(cfg *config.Configuration) *Server {
	return &Server{
		Router: echo.New(),
		Config: cfg,
	}
}

func (s *Server) Start() {
	s.Router.Logger.Fatal(
		s.Router.Start(":" + s.Config.Server.Port))
}

func (s *Server) InitRouter() {
	s.Router.Validator = utils.NewRequestValidator(validator.New())
	s.Router.Use(middleware.Logger()) //to log request info
	s.Router.Use(middleware.CORS())   //to allow cross-domain access
	s.Router.GET("/", indexHandler)
}

func indexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK,
		`Go Web Service is running. <a href="swagger/index.html">[See API definition]</a>`,
	)
}

func (s *Server) SetupSwagger(path string) {
	if s.Config.Swagger.Host == "" {
		s.Config.Swagger.Host = s.Config.Server.Host + ":" + s.Config.Server.Port
	}
	docs.SwaggerInfo.Host = s.Config.Swagger.Host
	s.Router.GET(path, echoSwagger.WrapHandler)
}

func (s *Server) RegisterRecordController(ctrl *record.Controller) {
	s.Router.POST("/records", ctrl.Create)
	s.Router.PUT("/records/:id", ctrl.Update)
	s.Router.DELETE("/records/:id", ctrl.Delete)
	s.Router.GET("/records/:id", ctrl.Get)
	s.Router.GET("/records", ctrl.Query)
}
