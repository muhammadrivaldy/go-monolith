package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"context"
	"time"

	"gorm.io/gorm"
)

func (s *securityUseCase) RegisterApi(ctx context.Context, req *payload.RequestRegisterApi) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: RegisterApi")
	defer span.End()

	// get info api by name
	resApi, err := s.securityEntity.ApiRepo.SelectApiByID(ctx, req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return
	}

	// if result id is zero, insert the api
	if resApi.Id == "" {

		// register endpoint with service id
		resApi, err = s.securityEntity.ApiRepo.InsertApi(models.Api{
			Id:        req.ID,
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

		// update detail api if the informations are different
		if resApi.Name != req.Name || resApi.Endpoint != req.Endpoint || resApi.Method != req.Method || resApi.ServiceID != req.ServiceID {

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

	}
}
