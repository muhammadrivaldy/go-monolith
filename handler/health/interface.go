package health

import (
	"backend/models"
	"context"
	"time"
)

// Usecase is a interface for layer business
type Usecase interface {
	HealthService(ctx context.Context) (res string, errs models.Error)
	HealthDB(ctx context.Context) (res time.Time, errs models.Error)
}

type Entity struct {
	Database
}

type Database interface {
	SelectTime() (res time.Time, err error)
}
