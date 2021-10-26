package usecase

import (
	"backend/models"
	"context"
	"net/http"
	"time"
)

func (u *usecase) RegisterService(ctx context.Context, serviceName string) (id int, errs models.Error) {

	// register service name
	res, err := u.securityEnti.InsertService(models.Service{
		Name:      serviceName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()})
	if err != nil {
		u.logs.Error(ctx, http.StatusInternalServerError, err)
		return id, models.ErrorMapping(err)
	}

	// if result id is zero, means it's a duplicate value.
	// get the result with servicename
	if res.ID == 0 {
		res, err = u.securityEnti.SelectServiceByName(serviceName)
		if err != nil {
			u.logs.Error(ctx, http.StatusInternalServerError, err)
			return id, models.ErrorMapping(err)
		}
	}

	// send result id of service name
	return res.ID, models.Error{}
}
