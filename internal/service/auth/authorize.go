package auth

import (
	"context"
	"errors"
	"fmt"

	errsdomain "github.com/Anton-Kraev/avito-test-assignment/internal/domain/errors"
	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/hash"
	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/tokens"
	"github.com/Anton-Kraev/avito-test-assignment/internal/repository"
)

const initBalance = 200

func (s *Service) Authorize(ctx context.Context, user models.User) (string, error) {
	const op = "auth.Service.Authorize"

	savedUser, err := s.userRepo.GetByName(ctx, user.Name)
	if errors.Is(err, repository.ErrNotFound) {
		user.Balance = initBalance
		err = s.userRepo.Insert(ctx, user)
	}

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if savedUser.Password != hash.GetPasswordHash(user.Password) {
		return "", fmt.Errorf("%s: %w", op, errsdomain.ErrInvalidPassword)
	}

	return tokens.CreateJWTToken(user)
}
