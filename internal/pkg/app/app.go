package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"helloWorld/internal/app/endpoint"
	"helloWorld/internal/app/middleware"
	"helloWorld/internal/app/service"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	app.service = service.New()
	app.endpoint = endpoint.New(app.service)

	app.echo = echo.New()

	app.echo.Use(middleware.RoleCheck)

	app.echo.GET("/status", app.endpoint.Status)

	return app, nil
}

func (app App) Run() error {
	log.Info().Msg("Starting server...")

	if err := app.echo.Start(":8000"); err != nil {
		log.Fatal().Err(err)
		return err
	}

	return nil
}
