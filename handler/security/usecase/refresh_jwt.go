package usecase

import (
	"backend/handler/security/payload"
	"backend/handler/users/models"
	"backend/logs"
	"backend/tracer"
	"backend/util"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

func (s *securityUseCase) RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error) {

	ctx, span := tracer.Tracer.Start(ctx, "UseCase: RefreshJWT")
	defer span.End()

	userInfo := util.GetContext(ctx)

	modelUser, err := s.userEntity.UserRepo.SelectUserByID(ctx, userInfo.UserID)
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

	res.UserID = modelUser.ID

	res.Token, err = createToken(modelUser, s.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	res.RefreshToken, err = createRefreshToken(modelUser, s.config.JWTKey)
	if err != nil {
		logs.Logging.Error(ctx, err)
		return res, util.ErrorMapping(err)
	}

	return res, util.ErrorMapping(nil)
}

func createRefreshToken(modelUser models.User, jwtKey string) (string, error) {

	requestCreateJWT := goutil.RequestCreateJWT{
		SignMethod: jwt.SigningMethodHS256,
		Key:        jwtKey,
		Data: jwt.MapClaims{
			"user_id":   modelUser.ID,
			"name":      modelUser.Name,
			"email":     modelUser.Email,
			"exp":       time.Now().AddDate(0, 0, 30).Unix(),
			"user_type": modelUser.UserTypeID,
			"type":      "refresh-token",
		},
	}

	token, err := goutil.CreateJWT(requestCreateJWT)
	if err != nil {
		return "", err
	}

	return token, err
}
