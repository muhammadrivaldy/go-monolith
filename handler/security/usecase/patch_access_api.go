package usecase

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
)

func (s *securityUseCase) PatchAccessApi(ctx context.Context, req payload.RequestPatchAccessApi) (errs util.Error) {

	// get info
	userInfo := util.GetContext(ctx)

	// validate user type
	if _, err := s.userEntity.UserTypeRepo.SelectUserTypeById(req.UserType); err == gorm.ErrRecordNotFound {
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

	apiId := []int{}
	for _, i := range access {
		apiId = append(apiId, i.ApiId)
	}

	sort.Ints(apiId)
	sort.Ints(req.ApiId)

	logs.Logging.Info(ctx, fmt.Sprintf("%+v", req.ApiId))

	apiIdToBeInsert := []int{}
	apiIdToBeDelete := []int{}

	if len(apiId) == 0 {

		apiIdToBeInsert = append(apiIdToBeInsert, req.ApiId...)

	} else if len(req.ApiId) == 0 {

		apiIdToBeDelete = append(apiIdToBeDelete, apiId...)

	} else {

		for i := 0; i < len(apiId); i++ {
			idx, _, exists := util.SearchIntInArray(apiId[i], req.ApiId)
			if exists {
				temp := req.ApiId[:idx]
				temp = append(temp, req.ApiId[idx+1:]...)
				req.ApiId = temp
			} else {
				apiIdToBeDelete = append(apiIdToBeDelete, apiId[i])
				temp := apiId[:i]
				temp = append(temp, apiId[i+1:]...)
				apiId = temp
				i--
			}

			if (len(apiId) - 1) == i {
				if len(req.ApiId) > 0 {
					apiIdToBeInsert = append(apiIdToBeInsert, req.ApiId...)
				}
			}
		}

	}

	// get time now
	timeNow := time.Now()

	// insert access if exists
	if len(apiIdToBeInsert) > 0 {
		accesses := []models.Access{}
		for _, i := range apiIdToBeInsert {
			accesses = append(accesses, models.Access{
				UserTypeId: req.UserType,
				ApiId:      i,
				CreatedBy:  userInfo.UserId,
				CreatedAt:  timeNow,
			})

			if _, err = s.securityEntity.AccessRepo.InsertAccesses(accesses); err != nil {
				logs.Logging.Error(ctx, err)
				return util.ErrorMapping(err)
			}
		}
	}

	// delete access if exists
	if len(apiIdToBeDelete) > 0 {
		if err = s.securityEntity.AccessRepo.DeleteAccessesByUserTypeIdAndApiId(req.UserType, apiIdToBeDelete); err != nil {
			logs.Logging.Error(ctx, err)
			return util.ErrorMapping(err)
		}
	}

	return

}
