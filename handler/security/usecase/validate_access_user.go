package usecase

import (
	"backend/logs"
	"backend/util"
	"context"

	"gorm.io/gorm"
)

func (u useCase) ValidateAccessUser(ctx context.Context, apiID int, userTypeID int) (bool, error) {

	// prepare a filter
	filter := util.FilterQuery{}
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "api_id", Operation: "=", Value: apiID})
	filter.Conditions = append(filter.Conditions, util.Condition{Operation: "and"})
	filter.Conditions = append(filter.Conditions, util.Condition{Field: "user_type_id", Operation: "=", Value: userTypeID})

	// check access of user
	_, err := u.securityEntity.AccessRepo.SelectAccessByFilter(filter)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return false, err
	}

	return true, nil
}
