package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DayLeft() int64
}

type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{
		service: service,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	days := e.service.DayLeft()

	s := fmt.Sprintf("Days left: %d", days)

	if err := ctx.String(http.StatusOK, s); err != nil {
		return err
	}

	return nil
}
