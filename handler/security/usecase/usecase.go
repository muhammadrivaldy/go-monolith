package usecase

import (
	"backend/config"
	"backend/handler/security"

	goutil "github.com/muhammadrivaldy/go-util"
)

type useCase struct {
	config       config.Configuration
	securityEnti security.Entity
	logs         goutil.Logs
}

// NewuseCase is a function for override interface
func NewuseCase(
	logs goutil.Logs,
	config config.Configuration,
	securityEnti security.Entity) security.UseCase {
	return &useCase{
		logs:         logs,
		config:       config,
		securityEnti: securityEnti}
}
