package usecase

import (
	"backend/handler/users/models"
	"backend/logs"
	"backend/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u userUseCase) GetUserById(ctx context.Context, id int64) (res models.User, errs util.Error) {

	user, err := u.userEntity.UserRepo.SelectUserById(id)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	if !user.Status.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	}

	return user, util.ErrorMapping(nil)
}
