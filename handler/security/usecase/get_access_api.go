package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"
	"sort"

	"gorm.io/gorm"
)

func (s *securityUseCase) GetAccessApi(ctx context.Context, req payload.RequestGetAccessApi) (res payload.ResponseGetAccessApi, errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: GetAccessApi")
	defer span.End()

	// validate user type
	_, err := s.userEntity.UserTypeRepo.SelectUserTypeByID(ctx, req.UserType)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	// get access by user type
	access, err := s.securityEntity.AccessRepo.SelectAccessByUserType(ctx, req.UserType)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	// mapping api id to response
	for _, i := range access {
		res.ApiID = append(res.ApiID, int(i.ApiID))
	}

	// sorting elements
	sort.Ints(res.ApiID)

	// response
	return

}
