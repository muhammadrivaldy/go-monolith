package usecase

import (
	"backend/handler/security"
	"backend/models"

	goutil "github.com/muhammadrivaldy/go-util"
)

type usecase struct {
	config       models.Configuration
	securityEnti security.Entity
	logs         goutil.Logs
}

// NewUsecase is a function for override interface
func NewUsecase(
	logs goutil.Logs,
	config models.Configuration,
	securityEnti security.Entity) security.Usecase {
	return &usecase{
		logs:         logs,
		config:       config,
		securityEnti: securityEnti}
}
