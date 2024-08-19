package usecase

import (
	"backend/handler/users/payload"
	"backend/logs"
	"backend/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u userUseCase) GetUserByID(ctx context.Context, req payload.RequestGetUserByID) (res payload.ResponseGetUserByID, errs util.Error) {

	// get user
	user, err := u.userEntity.UserRepo.SelectUserByID(req.UserID)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	if !user.Status.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	}

	// set response
	return payload.ResponseGetUserByID{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.Phone,
		Email:        user.Email,
		Status:       int(user.Status),
		StatusName:   user.Status.String(),
		UserType:     int(user.UserTypeID),
		UserTypeName: user.UserTypeID.String(),
	}, util.ErrorMapping(nil)
}
