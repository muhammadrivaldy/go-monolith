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

func (s securityUseCase) Login(ctx context.Context, email, password string) (res payload.ResponseLogin, errs util.Error) {

	modelUser, err := s.userEntity.UserRepo.SelectUserByEmail(email)
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

	if !modelUser.ValidatePassword(password) {
		logs.Logging.Error(ctx, errors.New("password is not match"))
		return res, util.ErrorMapping(util.ErrorUnauthorized)
	}

	res.Token, res.RefreshToken, err = goutil.CreateJWT(goutil.JWT{
		UserId:     modelUser.Id,
		Name:       modelUser.Name,
		Phone:      modelUser.Phone,
		Email:      modelUser.Email,
		GroupId:    int(modelUser.UserTypeId),
		ExpToken:   time.Now().Add(15 * time.Minute),
		ExpRefresh: time.Now().AddDate(0, 0, 15),
	}, jwt.SigningMethodHS256, s.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return res, util.ErrorMapping(nil)
}
