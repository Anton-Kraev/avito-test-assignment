package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"

	authmw "github.com/Anton-Kraev/avito-test-assignment/internal/http/middleware"
)

func New(log *slog.Logger) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(slogecho.NewWithConfig(log, slogecho.Config{
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  true,
		WithResponseBody:   true,
		WithResponseHeader: true,
	}))
	e.Use(authmw.JWTMiddlewareExcludeAuth)

	return e
}
