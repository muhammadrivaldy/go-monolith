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

func (s securityUseCase) RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error) {

	userInfo := goutil.GetContext(ctx)

	modelUser, err := s.userEntity.UserRepo.SelectUserById(userInfo.UserId)
	if err == gorm.ErrRecordNotFound {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	if !modelUser.Status.IsActive() {
		logs.Logging.Error(ctx, errors.New("user is not active"))
		return res, util.ErrorMapping(util.ErrorDataNotFound)
	}

	res.Token, res.RefreshToken, err = goutil.CreateJWT(goutil.JWT{
		UserId:     modelUser.Id,
		Name:       modelUser.Name,
		Phone:      modelUser.Phone,
		Email:      modelUser.Email,
		GroupId:    int(modelUser.UserTypeId),
		ExpToken:   time.Now().Add(15 * time.Minute),
		ExpRefresh: time.Now().AddDate(0, 0, 15),
		Jti:        userInfo.Jti,
	}, jwt.SigningMethodHS256, s.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return res, util.ErrorMapping(nil)
}
