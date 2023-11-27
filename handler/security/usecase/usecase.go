package usecase

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/entity"
	entityUser "backend/handler/users/entity"
)

type securityUseCase struct {
	config         config.Configuration
	securityEntity entity.SecurityEntity
	userEntity     entityUser.UserEntity
}

// NewUseCase is a function for override interface
func NewUseCase(
	config config.Configuration,
	securityEntity entity.SecurityEntity,
	userEntity entityUser.UserEntity) security.ISecurityUseCase {
	return &securityUseCase{
		config:         config,
		securityEntity: securityEntity,
		userEntity:     userEntity}
}
