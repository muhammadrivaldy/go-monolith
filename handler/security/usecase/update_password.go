package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (s securityUseCase) UpdatePassword(ctx context.Context, req payload.RequestUpdatePassword) (errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	// get user
	user, err := s.userEntity.UserRepo.SelectUserById(int64(req.UserId))
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	// generate password
	user.Password, err = util.GeneratePassword(req.Password)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	user.UpdatedBy = userInfo.UserId
	user.UpdatedAt = time.Now()

	// update the user password
	if _, err := s.userEntity.UserRepo.UpdateUser(user); err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	return

}
