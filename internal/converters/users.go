package converters

import (
	"github.com/ZetNetwork/Users/internal/domain/models/dto"
	"github.com/ZetNetwork/Users/internal/domain/models/entities"
)

type UsersConverter struct{}

func NewUsersConverter() *UsersConverter {
	return &UsersConverter{}
}

func (u UsersConverter) UserToDTO(user entities.User) dto.User {
	return dto.User{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Surname:  user.Surname,
	}
}

func (u UsersConverter) UserToEntity(user dto.User) entities.User {
	return entities.User{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Surname:  user.Surname,
	}
}
