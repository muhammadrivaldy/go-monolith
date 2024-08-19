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
	if _, err := s.userEntity.UserTypeRepo.SelectUserTypeByID(req.UserType); err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	// get access by user type
	access, err := s.securityEntity.AccessRepo.SelectAccessByUserType(req.UserType)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	apiID := []int{}
	for _, i := range access {
		apiID = append(apiID, i.ApiID)
	}

	sort.Ints(apiID)
	sort.Ints(req.ApiID)

	logs.Logging.Info(ctx, fmt.Sprintf("%+v", req.ApiID))

	apiIDToBeInsert := []int{}
	apiIDToBeDelete := []int{}

	if len(apiID) == 0 {

		apiIDToBeInsert = append(apiIDToBeInsert, req.ApiID...)

	} else if len(req.ApiID) == 0 {

		apiIDToBeDelete = append(apiIDToBeDelete, apiID...)

	} else {

		for i := 0; i < len(apiID); i++ {
			idx, _, exists := util.SearchIntInArray(apiID[i], req.ApiID)
			if exists {
				temp := req.ApiID[:idx]
				temp = append(temp, req.ApiID[idx+1:]...)
				req.ApiID = temp
			} else {
				apiIDToBeDelete = append(apiIDToBeDelete, apiID[i])
				temp := apiID[:i]
				temp = append(temp, apiID[i+1:]...)
				apiID = temp
				i--
			}

			if (len(apiID) - 1) == i {
				if len(req.ApiID) > 0 {
					apiIDToBeInsert = append(apiIDToBeInsert, req.ApiID...)
				}
			}
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

			if _, err = s.securityEntity.AccessRepo.InsertAccesses(accesses); err != nil {
				logs.Logging.Error(ctx, err)
				return util.ErrorMapping(err)
			}
		}
	}

	// delete access if exists
	if len(apiIDToBeDelete) > 0 {
		if err = s.securityEntity.AccessRepo.DeleteAccessesByUserTypeIDAndApiID(req.UserType, apiIDToBeDelete); err != nil {
			logs.Logging.Error(ctx, err)
			return util.ErrorMapping(err)
		}
	}

	return

}
