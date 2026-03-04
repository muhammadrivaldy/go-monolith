package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
)

func (s *securityUseCase) PatchAccessApi(ctx context.Context, req payload.RequestPatchAccessApi) (errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: PatchAccessApi")
	defer span.End()

	// get info
	userInfo := util.GetContext(ctx)

	// validate user type
	if _, err := s.userEntity.UserTypeRepo.SelectUserTypeByID(ctx, req.UserType); err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	// get access by user type
	access, err := s.securityEntity.AccessRepo.SelectAccessByUserType(ctx, req.UserType)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	apiID := []string{}
	for _, i := range access {
		apiID = append(apiID, i.ApiID)
	}

	sort.Strings(apiID)
	sort.Strings(req.ApiID)

	logs.Logging.Info(ctx, fmt.Sprintf("%+v", req.ApiID))

	apiIDToBeInsert := []string{}
	apiIDToBeDelete := []string{}

	currentAccessMap := map[string]struct{}{}
	for _, id := range apiID {
		currentAccessMap[id] = struct{}{}
	}

	requestedAccessMap := map[string]struct{}{}
	for _, id := range req.ApiID {
		requestedAccessMap[id] = struct{}{}
	}

	for id := range requestedAccessMap {
		if _, exists := currentAccessMap[id]; !exists {
			apiIDToBeInsert = append(apiIDToBeInsert, id)
		}
	}

	for id := range currentAccessMap {
		if _, exists := requestedAccessMap[id]; !exists {
			apiIDToBeDelete = append(apiIDToBeDelete, id)
		}
	}

	// get time now
	timeNow := time.Now()

	// insert access if exists
	if len(apiIDToBeInsert) > 0 {
		accesses := []models.Access{}
		for _, i := range apiIDToBeInsert {
			accesses = append(accesses, models.Access{
				UserTypeID: req.UserType,
				ApiID:      i,
				CreatedBy:  userInfo.UserID,
				CreatedAt:  timeNow,
			})
		}

		if _, err = s.securityEntity.AccessRepo.InsertAccesses(ctx, accesses); err != nil {
			logs.Logging.Error(ctx, err)
			return util.ErrorMapping(err)
		}
	}

	// delete access if exists
	if len(apiIDToBeDelete) > 0 {
		if err = s.securityEntity.AccessRepo.DeleteAccessesByUserTypeIDAndApiID(ctx, req.UserType, apiIDToBeDelete); err != nil {
			logs.Logging.Error(ctx, err)
			return util.ErrorMapping(err)
		}
	}

	return

}
