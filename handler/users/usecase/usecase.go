package usecase

import (
	"backend/config"
	"backend/handler/users"
	"backend/handler/users/entity"
)

type userUseCase struct {
	config     config.Configuration
	userEntity entity.UserEntity
}

// NewUseCase is a function for override interface
func NewUseCase(
	config config.Configuration,
	userEntity entity.UserEntity) users.IUserUseCase {
	return userUseCase{
		config:     config,
		userEntity: userEntity}
}
