package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"
	"time"

	"gorm.io/gorm"
)

func (s securityUseCase) EditPassword(ctx context.Context, req payload.RequestEditPassword) (errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: EditPassword")
	defer span.End()

	userInfo := util.GetContext(ctx)

	// get user
	user, err := s.userEntity.UserRepo.SelectUserByID(ctx, req.UserID)
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

	user.UpdatedBy = userInfo.UserID
	user.UpdatedAt = time.Now()

	// update the user password
	if _, err := s.userEntity.UserRepo.UpdateUser(ctx, user); err != nil {
		logs.Logging.Error(ctx, err)
		return util.ErrorMapping(err)
	}

	return

}
