package handler

import (
	"context"
	"log/slog"

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
