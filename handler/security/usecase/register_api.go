package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/logs"
	"context"
	"time"

	"gorm.io/gorm"
)

func (s *securityUseCase) RegisterApi(ctx context.Context, req *payload.RequestRegisterApi) {

	// get info api by name
	resApi, err := s.securityEntity.ApiRepo.SelectApiByName(req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return
	}

	// get info api by endpoint
	if resApi.Id == 0 {
		resApi, err = s.securityEntity.ApiRepo.SelectApiByEndpoint(req.Endpoint, req.Method)
		if err != nil && err != gorm.ErrRecordNotFound {
			logs.Logging.Error(ctx, err)
			return
		}
	}

	// if result id is zero, insert the api
	if resApi.Id == 0 {

		// register endpoint with service id
		resApi, err = s.securityEntity.ApiRepo.InsertApi(models.Api{
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
		if err != nil {
			logs.Logging.Error(ctx, err)
			return
		}

	} else {

		// update detail api
		if _, err = s.securityEntity.ApiRepo.UpdateApi(models.Api{
			Id:        resApi.Id,
			Name:      req.Name,
			Endpoint:  req.Endpoint,
			Method:    req.Method,
			ServiceID: req.ServiceID}); err != nil {
			logs.Logging.Error(ctx, err)
			return
		}

	}

	// override the id of api
	req.ID = resApi.Id
}
