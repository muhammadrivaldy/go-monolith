package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"

	"gorm.io/gorm"
)

func (s *securityUseCase) GetApisByServiceID(ctx context.Context, req payload.RequestGetApisServiceID) (res []payload.ResponseGetApisServiceID, errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: GetApisByServiceID")
	defer span.End()

	// validate service id
	_, err := s.securityEntity.ServiceRepo.SelectServiceByID(ctx, req.ServiceID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	// get apis
	apis, err := s.securityEntity.ApiRepo.SelectApisByServiceID(ctx, req.ServiceID)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	for _, i := range apis {
		res = append(res, payload.ResponseGetApisServiceID{
			ApiID:    i.ID,
			ApiName:  i.Name,
			Method:   i.Method,
			Endpoint: i.Endpoint,
		})
	}

	return

}
