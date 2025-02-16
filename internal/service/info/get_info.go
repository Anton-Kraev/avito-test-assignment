package info

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

func (s *Service) GetInfo(ctx context.Context, userID int) (models.Info, error) {
	const op = "info.Service.GetInfo"

	balance, err := s.userRepo.GetBalanceByID(ctx, userID)
	if err != nil {
		return models.Info{}, fmt.Errorf("%s: %w", op, err)
	}

	receivedTransfers, err := s.transferRepo.GetAllByReceiverUserID(ctx, userID)
	if err != nil {
		return models.Info{}, fmt.Errorf("%s: %w", op, err)
	}

	sentTransfers, err := s.transferRepo.GetAllBySenderUserID(ctx, userID)
	if err != nil {
		return models.Info{}, fmt.Errorf("%s: %w", op, err)
	}

	purchases, err := s.purchaseRepo.GetAllByUserID(ctx, userID)
	if err != nil {
		return models.Info{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.Info{
		Balance:             balance,
		Inventory:           models.InventoryFromMerchList(purchases),
		CoinsReceiveHistory: receivedTransfers,
		CoinsSentHistory:    sentTransfers,
	}, nil
}
