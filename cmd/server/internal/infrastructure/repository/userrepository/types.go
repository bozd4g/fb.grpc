package userrepository

import (
	"github.com/bozd4g/fb.grpc/cmd/server/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Add(entity user.Entity) error
	Get(id uuid.UUID) (*user.Entity, error)
	GetAll() ([]user.Entity, error)
}

type UserRepository struct {
	db *gorm.DB
}
