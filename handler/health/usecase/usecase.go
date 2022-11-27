package usecase

import (
	health "backend/handler/health"

	goutil "github.com/muhammadrivaldy/go-util"
)

type useCase struct {
	logs   goutil.Logs
	health health.Entity
}

// NewuseCase is a function for override interface
func NewuseCase(
	logs goutil.Logs,
	healthEnti health.Entity) health.UseCase {
	return &useCase{
		logs:   logs,
		health: healthEnti}
}
