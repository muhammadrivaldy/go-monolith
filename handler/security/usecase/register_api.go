package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/logs"
	"context"
	"time"

	"gorm.io/gorm"
)

func (u *useCase) RegisterApi(req *payload.RegisterApiRequest) {

	// get info api by name
	resApi, err := u.securityEntity.ApiRepo.SelectApiByName(req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Logging.Error(context.Background(), err)
		return
	}

	// get info api by endpoint
	if resApi.ID == 0 {
		resApi, err = u.securityEntity.ApiRepo.SelectApiByEndpoint(req.Endpoint)
		if err != nil && err != gorm.ErrRecordNotFound {
			logs.Logging.Error(context.Background(), err)
			return
		}
	}

	// if result id is zero, insert the api
	if resApi.ID == 0 {

		// register endpoint with service id
		resApi, err = u.securityEntity.ApiRepo.InsertApi(models.Api{
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
		if err != nil {
			logs.Logging.Error(context.Background(), err)
			return
		}
	} else {

		// update detail api
		if _, err = u.securityEntity.ApiRepo.UpdateApi(models.Api{
			ID:        resApi.ID,
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID}); err != nil {
			logs.Logging.Error(context.Background(), err)
			return
		}
	}

	// override the id of api
	req.ID = resApi.ID
}
