package usecase

import (
	"backend/util"
	"context"
	"time"
)

func (u *usecase) HealthDB(ctx context.Context) (res time.Time, errs util.Error) {
	res, err := u.health.SelectTime()
	if err != nil {
		u.logs.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return
}
