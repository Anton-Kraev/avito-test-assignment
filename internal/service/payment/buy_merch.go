package payment

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/errors"
	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

func (s *Service) BuyMerch(ctx context.Context, userID int, merchName string) error {
	const op = "payment.Service.BuyMerch"

	balance, err := s.userRepo.GetBalanceByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	merch, err := s.merchRepo.GetByName(merchName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if merch.Price > balance {
		return fmt.Errorf("%s: %w", op, errors.ErrNotEnoughMoney)
	}

	if err = s.userRepo.UpdateBalanceByID(ctx, userID, -merch.Price); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.purchaseRepo.Insert(ctx, models.Purchase{
		UserID:    userID,
		MerchType: merchName,
	}); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
