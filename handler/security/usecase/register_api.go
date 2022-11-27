package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"context"
	"time"

	"gorm.io/gorm"
)

func (u *useCase) RegisterAPI(req *payload.RegisterAPIRequest) {

	// get info api by name
	resAPI, err := u.securityEnti.SelectAPIByName(req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		u.logs.Error(context.Background(), err)
		return
	}

	// get info api by endpoint
	if resAPI.ID == 0 {
		resAPI, err = u.securityEnti.SelectAPIByEndpoint(req.Endpoint)
		if err != nil && err != gorm.ErrRecordNotFound {
			u.logs.Error(context.Background(), err)
			return
		}
	}

	// if result id is zero, insert the api
	if resAPI.ID == 0 {

		// register endpoint with service id
		resAPI, err = u.securityEnti.InsertAPI(models.API{
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
		if err != nil {
			u.logs.Error(context.Background(), err)
			return
		}
	} else {

		// update detail api
		if _, err = u.securityEnti.UpdateAPI(models.API{
			ID:        resAPI.ID,
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID}); err != nil {
			u.logs.Error(context.Background(), err)
			return
		}
	}

	// override the id of api
	req.ID = resAPI.ID
}
