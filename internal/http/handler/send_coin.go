package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
	"github.com/Anton-Kraev/avito-test-assignment/internal/logger"
	"github.com/Anton-Kraev/avito-test-assignment/openapi/api"
)

func (h Handler) PostApiSendCoin(ctx echo.Context) error {
	const op = "handler.Handler.PostApiSendCoin"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
	)

	var req api.PostApiSendCoinJSONRequestBody

	if err := ctx.Bind(&req); err != nil {
		errMsg := "invalid request"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Errors: &errMsg})
	}

	senderID := ctx.Get("userID").(int)

	err := h.paymentService.SendCoins(ctx.Request().Context(), models.Transfer{
		ReceiverName: req.ToUser,
		SenderID:     senderID,
		Amount:       req.Amount,
	})
	if err != nil {
		errMsg := "send coins failed"

		log.Error(errMsg, logger.Err(err))

		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Errors: &errMsg})
	}

	log.Info("successfully sent coins")

	return ctx.String(http.StatusOK, "successfully sent coins")
}
