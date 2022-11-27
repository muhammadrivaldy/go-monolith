package usecase

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/entity"
)

type useCase struct {
	config         config.Configuration
	securityEntity entity.SecurityEntity
}

// NewSecurityUseCase is a function for override interface
func NewSecurityUseCase(
	config config.Configuration,
	securityEntity entity.SecurityEntity) security.ISecurityUseCase {
	return &useCase{
		config:         config,
		securityEntity: securityEntity}
}
