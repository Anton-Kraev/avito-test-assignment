package payment

import (
	"context"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

type (
	userRepo interface {
		GetBalanceByID(ctx context.Context, id int) (int, error)
		UpdateBalanceByID(ctx context.Context, userID int, amount int) error
	}

	merchRepo interface {
		GetByName(name string) (models.Merch, error)
	}

	purchaseRepo interface {
		Insert(ctx context.Context, purchase models.Purchase) error
	}

	transferRepo interface {
		Insert(ctx context.Context, transfer models.Transfer) error
	}
)

type Service struct {
	userRepo     userRepo
	merchRepo    merchRepo
	purchaseRepo purchaseRepo
	transferRepo transferRepo
}

func NewService(userRepo userRepo, merchRepo merchRepo, purchaseRepo purchaseRepo, transferRepo transferRepo) *Service {
	return &Service{
		userRepo:     userRepo,
		merchRepo:    merchRepo,
		purchaseRepo: purchaseRepo,
		transferRepo: transferRepo,
	}
}
