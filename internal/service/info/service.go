package info

import (
	"context"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

type (
	userRepo interface {
		GetBalanceByID(ctx context.Context, id int) (int, error)
	}

	purchaseRepo interface {
		GetAllByUserID(ctx context.Context, id int) ([]string, error)
	}

	transferRepo interface {
		GetAllByReceiverUserID(ctx context.Context, id int) ([]models.ReceivedTransfer, error)
		GetAllBySenderUserID(ctx context.Context, id int) ([]models.SentTransfer, error)
	}
)

type Service struct {
	userRepo     userRepo
	purchaseRepo purchaseRepo
	transferRepo transferRepo
}

func NewService(userRepo userRepo, purchaseRepo purchaseRepo, transferRepo transferRepo) *Service {
	return &Service{
		userRepo:     userRepo,
		purchaseRepo: purchaseRepo,
		transferRepo: transferRepo,
	}
}
