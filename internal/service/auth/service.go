package auth

import (
	"context"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

type userRepo interface {
	GetByName(ctx context.Context, name string) (models.User, error)
	Insert(ctx context.Context, user models.User) error
}

type Service struct {
	userRepo userRepo
}

func NewService(userRepo userRepo) *Service {
	return &Service{userRepo: userRepo}
}
