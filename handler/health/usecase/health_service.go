package usecase

import (
	"backend/util"
	"context"
)

func (u *useCase) HealthService(ctx context.Context) (res string, errs util.Error) {
	return "Success", errs
}
