package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/internal/entity"
)

type AuthRepository interface {
	Register(ctx context.Context, user *entity.Users) error
	Login(ctx context.Context, email, password string) (*entity.Users, error)
}
