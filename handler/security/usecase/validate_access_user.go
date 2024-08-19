package usecase

import (
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"

	"gorm.io/gorm"
)

func (s *securityUseCase) ValidateAccessUser(ctx context.Context, apiID string) (res bool, errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: ValidateAccessUser")
	defer span.End()

	userInfo := util.GetContext(ctx)

	// prepare a filter
	filter := util.FilterQuery{}
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "api_id", Operation: "=", Value: apiID})
	filter.Conditions = append(filter.Conditions, util.Condition{Operation: "and"})
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "user_type_id", Operation: "=", Value: userInfo.UserType})

	// check access of user
	_, err := s.securityEntity.AccessRepo.SelectAccessByFilter(ctx, filter)
	if err == gorm.ErrRecordNotFound {
		return false, util.ErrorMapping(util.ErrorUserDoesNotHaveAuthorization)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return false, util.ErrorMapping(err)
	}

	return true, util.ErrorMapping(nil)
}
