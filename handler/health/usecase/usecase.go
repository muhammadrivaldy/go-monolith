package usecase

import (
	health "backend/handler/health"

	goutil "github.com/muhammadrivaldy/go-util"
)

type usecase struct {
	logs   goutil.Logs
	health health.Entity
}

// NewUsecase is a function for override interface
func NewUsecase(
	logs goutil.Logs,
	healthEnti health.Entity) health.Usecase {
	return &usecase{
		logs:   logs,
		health: healthEnti}
}
