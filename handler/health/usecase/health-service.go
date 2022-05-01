package usecase

import (
	"backend/util"
	"context"
)

func (u *usecase) HealthService(ctx context.Context) (res string, errs util.Error) {
	return "Success", errs
}
