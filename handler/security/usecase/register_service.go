package usecase

import (
	"backend/handler/security/models"
	"backend/logs"
	"backend/util"
	"context"
	"time"

	"gorm.io/gorm"
)

func (s *securityUseCase) RegisterService(ctx context.Context, serviceName string) (id int64, errs util.Error) {

	res, err := s.securityEntity.ServiceRepo.SelectServiceByName(serviceName)
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return id, util.ErrorMapping(err)
	}

	// register service name
	if res.ID == 0 {
		res, err = s.securityEntity.ServiceRepo.InsertService(models.Service{
			Name:      serviceName,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
		if err != nil {
			logs.Logging.Error(ctx, err)
			return id, util.ErrorMapping(err)
		}
	}

	// send result id of service name
	return res.ID, util.ErrorMapping(nil)
}
