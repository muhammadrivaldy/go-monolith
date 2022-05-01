package usecase

import (
	"backend/config"
	"backend/handler/security"

	goutil "github.com/muhammadrivaldy/go-util"
)

type usecase struct {
	config       config.Configuration
	securityEnti security.Entity
	logs         goutil.Logs
}

// NewUsecase is a function for override interface
func NewUsecase(
	logs goutil.Logs,
	config config.Configuration,
	securityEnti security.Entity) security.Usecase {
	return &usecase{
		logs:         logs,
		config:       config,
		securityEnti: securityEnti}
}
