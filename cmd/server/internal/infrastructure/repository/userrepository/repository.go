package userrepository

import (
	"errors"
	"github.com/bozd4g/fb.grpc/cmd/server/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func New(db *gorm.DB) IUserRepository {
	return UserRepository{db: db}
}

func (repository UserRepository) Add(entity user.Entity) error {
	transaction := repository.db.Begin()
	defer func() error {
		if r := recover(); r != nil {
			transaction.Rollback()
		}

		return errors.New("An error occured while creating a new user!")
	}()

	entity.Id = uuid.New()
	transaction.Create(&entity)
	if err := transaction.Error; err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (repository UserRepository) Get(id uuid.UUID) (*user.Entity, error) {
	var entity user.Entity
	result := repository.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository UserRepository) GetAll() ([]user.Entity, error) {
	var users []user.Entity
	result := repository.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
