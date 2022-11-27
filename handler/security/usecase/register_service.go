package usecase

import (
	"backend/handler/security/models"
	"backend/logs"
	"backend/util"
	"context"
	"time"
)

func (u *useCase) RegisterService(ctx context.Context, serviceName string) (id int, errs util.Error) {

	// register service name
	res, err := u.securityEntity.ServiceRepo.InsertService(models.Service{
		Name:      serviceName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()})
	if err != nil {
		logs.Logging.Error(ctx, err)
		return id, util.ErrorMapping(err)
	}

	// if result id is zero, means it's a duplicate value.
	// get the result with service name
	if res.ID == 0 {
		res, err = u.securityEntity.ServiceRepo.SelectServiceByName(serviceName)
		if err != nil {
			logs.Logging.Error(ctx, err)
			return id, util.ErrorMapping(err)
		}
	}

	// send result id of service name
	return res.ID, util.ErrorMapping(err)
}
