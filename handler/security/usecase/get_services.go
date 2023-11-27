package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"
)

func (s *securityUseCase) GetServices(ctx context.Context) (res []payload.ResponseGetServices, errs util.Error) {

	// get services
	services, err := s.securityEntity.ServiceRepo.SelectServices()
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	for _, i := range services {
		res = append(res, payload.ResponseGetServices{
			ServiceID:   int(i.Id),
			ServiceName: i.Name,
		})
	}

	return

}
