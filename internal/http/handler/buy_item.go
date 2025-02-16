package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/logger"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

func (h Handler) GetApiBuyItem(ctx echo.Context, item string) error {
	const op = "handler.Handler.GetApiBuyItem"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
	)

	userID := ctx.Get("userID").(int)

	if err := h.paymentService.BuyMerch(ctx.Request().Context(), userID, item); err != nil {
		errMsg := "buy merch failed"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Errors: &errMsg})
	}

	log.Info("successfully bought merch")

	return ctx.String(http.StatusOK, "successfully bought merch")
}
