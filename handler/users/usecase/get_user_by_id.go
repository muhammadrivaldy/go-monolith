package usecase

import (
	"backend/handler/users/payload"
	"backend/logs"
	"backend/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u userUseCase) GetUserById(ctx context.Context, req payload.RequestGetUserById) (res payload.ResponseGetUserById, errs util.Error) {

	// get user
	user, err := u.userEntity.UserRepo.SelectUserById(req.UserId)
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
	return payload.ResponseGetUserById{
		Id:           user.Id,
		Name:         user.Name,
		PhoneNumber:  user.Phone,
		Email:        user.Email,
		Status:       int(user.Status),
		StatusName:   user.Status.String(),
		UserType:     int(user.UserTypeId),
		UserTypeName: user.UserTypeId.String(),
	}, util.ErrorMapping(nil)
}
