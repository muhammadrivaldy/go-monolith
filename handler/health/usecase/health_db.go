package usecase

import (
	"backend/util"
	"context"
	"time"
)

func (u *useCase) HealthDB(ctx context.Context) (res time.Time, errs util.Error) {

	res, err := u.healthEntity.HealthRepo.SelectTime()
	if err != nil {
		u.logs.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return
}
