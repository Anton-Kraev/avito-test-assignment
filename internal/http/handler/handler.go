package handler

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

type (
	authService interface {
		Authorize(ctx context.Context, user models.User) (string, error)
	}

	infoService interface {
		GetInfo(ctx context.Context, userID int) (models.Info, error)
	}

	paymentService interface {
		BuyMerch(ctx context.Context, userID int, merchName string) error
		SendCoins(ctx context.Context, transfer models.Transfer) error
	}
)

type Handler struct {
	log            *slog.Logger
	authService    authService
	infoService    infoService
	paymentService paymentService
}

func New(log *slog.Logger, authService authService, infoService infoService, paymentService paymentService) Handler {
	return Handler{
		log:            log,
		authService:    authService,
		infoService:    infoService,
		paymentService: paymentService,
	}
}

func (h Handler) PostApiAuth(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetApiBuyItem(ctx echo.Context, item string) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetApiInfo(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostApiSendCoin(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
