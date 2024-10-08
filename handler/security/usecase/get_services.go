package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"
)

func (s *securityUseCase) GetServices(ctx context.Context) (res []payload.ResponseGetServices, errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: GetServices")
	defer span.End()

	// get services
	services, err := s.securityEntity.ServiceRepo.SelectServices(ctx)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	for _, i := range services {
		res = append(res, payload.ResponseGetServices{
			ServiceID:   int(i.ID),
			ServiceName: i.Name,
		})
	}

	return

}
