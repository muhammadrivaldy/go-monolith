package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (s *securityUseCase) RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser, err := s.userEntity.UserRepo.SelectUserById(userInfo.UserId)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Warning(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	if !modelUser.Status.IsActive() {
		logs.Logging.Warning(ctx, errors.New("user is not active"))
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	}

	res.UserId = modelUser.Id

	res.Token, res.RefreshToken, err = goutil.CreateJWT(goutil.JWT{
		UserId:     modelUser.Id,
		UserType:   int(modelUser.UserTypeId),
		Email:      modelUser.Email,
		ExpToken:   time.Now().AddDate(0, 0, 2),
		ExpRefresh: time.Now().AddDate(0, 0, 15),
		Jti:        userInfo.Jti,
	}, jwt.SigningMethodHS256, s.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return res, util.ErrorMapping(nil)
}
