package health

import (
	"backend/util"
	"context"
	"time"
)

// UseCase is a interface for layer business
type IHealthUseCase interface {
	HealthService(ctx context.Context) (res string, errs util.Error)
	HealthDB(ctx context.Context) (res time.Time, errs util.Error)
}

type IHealthRepo interface {
	SelectTime() (res time.Time, err error)
}
