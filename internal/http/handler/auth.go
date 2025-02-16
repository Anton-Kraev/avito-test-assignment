package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	errsdomain "github.com/Anton-Kraev/avito-test-assignment/internal/domain/errors"
	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/logger"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

func (h Handler) PostApiAuth(ctx echo.Context) error {
	const op = "handler.Handler.PostApiAuth"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
	)

	var req api.PostApiAuthJSONRequestBody

	if err := ctx.Bind(&req); err != nil {
		errMsg := "invalid request"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Errors: &errMsg})
	}

	token, err := h.authService.Authorize(ctx.Request().Context(), models.User{
		Name:     req.Username,
		Password: req.Password,
	})

	switch {
	case errors.Is(err, errsdomain.ErrInvalidPassword):
		if err != nil {
			errMsg := "invalid password"

			log.Error(errMsg, logger.Err(err))

			return ctx.JSON(http.StatusUnauthorized, api.ErrorResponse{Errors: &errMsg})
		}
	case err != nil:
		errMsg := "authentication failed"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Errors: &errMsg})
	}

	log.Info("successfully authenticated")

	return ctx.JSON(http.StatusOK, api.AuthResponse{Token: &token})
}
