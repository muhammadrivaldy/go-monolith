package usecase

import (
	"backend/models"
	"context"
	"net/http"
	"time"
)

func (u *usecase) HealthService(ctx context.Context) (res string, errs models.Error) {
	return "Success", errs
}

func (u *usecase) HealthDB(ctx context.Context) (res time.Time, errs models.Error) {
	res, err := u.health.SelectTime()
	if err != nil {
		u.logs.Error(ctx, http.StatusInternalServerError, err)
		return res, models.ErrorMapping(err)
	}

	return
}
