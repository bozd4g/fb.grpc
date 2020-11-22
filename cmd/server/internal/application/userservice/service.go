package userservice

import (
	"github.com/bozd4g/fb.grpc/cmd/server/internal/domain/user"
	"github.com/bozd4g/fb.grpc/cmd/server/internal/infrastructure/repository/userrepository"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func New(repository userrepository.IUserRepository) IUserService {
	return UserService{repository: repository}
}

func (service UserService) Create(userDto UserCreateRequestDto) error {
	var entity user.Entity
	err := mapstructure.Decode(userDto, &entity)
	if err != nil {
		return err
	}

	err = service.repository.Add(entity)
	if err != nil {
		return err
	}

	return nil
}

func (service UserService) GetAll() ([]UserDto, error) {
	dtos := make([]UserDto, 0)

	users, err := service.repository.GetAll()
	if err != nil {
		return dtos, err
	}

	for _, entity := range users {
		dtos = append(dtos, UserDto{
			Id:       entity.Id.String(),
			Name:     entity.Name,
			Surname:  entity.Surname,
			Email:    entity.Email,
			Password: entity.Password,
		})
	}

	return dtos, nil
}

func (service UserService) Get(id uuid.UUID) (*UserDto, error) {
	user, err := service.repository.Get(id)
	if err != nil {
		return nil, err
	}

	var dto UserDto
	err = mapstructure.Decode(user, &dto)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
