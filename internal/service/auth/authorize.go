package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
	"github.com/Anton-Kraev/avito-test-assignment/internal/repository"
)

func (s *Service) Authorize(ctx context.Context, user models.User) (string, error) {
	const op = "auth.Service.Authorize"

	err := s.userRepo.FindByName(ctx, user.Name)
	if errors.Is(err, repository.ErrNotFound) {
		user.Balance = initBalance
		err = s.userRepo.Insert(ctx, user)
	}

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return createToken()
}
