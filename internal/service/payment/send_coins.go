package payment

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/errors"
	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

func (s *Service) SendCoins(ctx context.Context, transfer models.Transfer) error {
	const op = "payment.Service.SendCoins"

	senderBalance, err := s.userRepo.GetBalanceByID(ctx, transfer.SenderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if transfer.Amount > senderBalance {
		return fmt.Errorf("%s: %w", op, errors.ErrNotEnoughMoney)
	}

	if err = s.userRepo.UpdateBalanceByID(ctx, transfer.SenderID, -transfer.Amount); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.userRepo.UpdateBalanceByID(ctx, transfer.ReceiverID, transfer.Amount); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.transferRepo.Insert(ctx, transfer); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
