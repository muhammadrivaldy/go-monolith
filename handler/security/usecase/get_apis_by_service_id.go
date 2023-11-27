package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"

	"gorm.io/gorm"
)

func (s *securityUseCase) GetApisByServiceId(ctx context.Context, req payload.RequestGetApisServiceId) (res []payload.ResponseGetApisServiceId, errs util.Error) {

	// validate service id
	_, err := s.securityEntity.ServiceRepo.SelectServiceById(req.ServiceID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	// get apis
	apis, err := s.securityEntity.ApiRepo.SelectApisByServiceId(req.ServiceID)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	for _, i := range apis {
		res = append(res, payload.ResponseGetApisServiceId{
			ApiId:    int(i.Id),
			ApiName:  i.Name,
			Method:   i.Method,
			Endpoint: i.Endpoint,
		})
	}

	return

}
