package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/logger"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

func (h Handler) GetApiInfo(ctx echo.Context) error {
	const op = "handler.Handler.GetApiInfo"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
	)

	userID := ctx.Get("userID").(int)

	info, err := h.infoService.GetInfo(ctx.Request().Context(), userID)
	if err != nil {
		errMsg := "get info failed"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Errors: &errMsg})
	}

	log.Info("successfully get info")

	return ctx.JSON(http.StatusOK, api.InfoResponse{Coins: &info.Balance})
}
