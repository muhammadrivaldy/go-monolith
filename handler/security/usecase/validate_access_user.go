package usecase

import (
	"backend/logs"
	"backend/util"
	"context"

	goutil "github.com/muhammadrivaldy/go-util"
)

func (s *securityUseCase) ValidateAccessUser(ctx context.Context, apiId int64) (res bool, errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	// prepare a filter
	filter := util.FilterQuery{}
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "api_id", Operation: "=", Value: apiId})
	filter.Conditions = append(filter.Conditions, util.Condition{Operation: "and"})
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "user_type_id", Operation: "=", Value: userInfo.GroupId})

	// check access of user
	_, err := s.securityEntity.AccessRepo.SelectAccessByFilter(filter)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return false, util.ErrorMapping(err)
	}

	return true, util.ErrorMapping(nil)
}
