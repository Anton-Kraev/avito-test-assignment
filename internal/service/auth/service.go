package auth

import (
	"context"

	"github.com/Anton-Kraev/avito-test-assignment/internal/domain/models"
)

type userRepo interface {
	FindByName(ctx context.Context, name string) error
	Insert(ctx context.Context, user models.User) error
}

type Service struct {
	userRepo userRepo
}

func NewService(userRepo userRepo) *Service {
	return &Service{userRepo: userRepo}
}

// TODO: get default balance from config
const initBalance = 200

// TODO: create jwt token
func createToken() (string, error) {
	return "", nil
}
