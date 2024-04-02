package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(value, roleAdmin) {
			log.Info().Msg("red button user detected")
		}

		if err := next(ctx); err != nil {
			return err
		}
		return nil
	}
}
