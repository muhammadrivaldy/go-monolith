package middleware

import (
	"backend/handler/security"
	"backend/handler/users"
	"backend/logs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type middleware struct {
	useCaseUser     users.IUserUseCase
	useCaseSecurity security.ISecurityUseCase
}

func NewMiddleware(useCaseUser users.IUserUseCase, useCaseSecurity security.ISecurityUseCase) middleware {
	return middleware{
		useCaseUser:     useCaseUser,
		useCaseSecurity: useCaseSecurity,
	}
}

func (m middleware) ValidateAccess(apiId int64) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := goutil.ParseContext(c)

		userInfo := goutil.GetContext(ctx)

		modelUser, errs := m.useCaseUser.GetUserById(ctx, userInfo.UserId)
		if errs.Error != nil {
			logs.Logging.Error(ctx, errs.Error)
			goutil.ResponseError(c, http.StatusForbidden, errors.New(http.StatusText(http.StatusForbidden)), nil)
			c.Abort()
			return
		}

		// the meaning of 2 is api for refresh jwt and make sure the id of refresh jwt is 2
		if apiId != 2 {
			if modelUser.Email != userInfo.Email || modelUser.Phone != userInfo.Phone || int(modelUser.UserTypeId) != userInfo.GroupId {
				logs.Logging.Error(ctx, errors.New("failed when validate the user data"))
				goutil.ResponseError(c, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)), nil)
				c.Abort()
				return
			}
		}

		_, errs = m.useCaseSecurity.ValidateAccessUser(ctx, apiId)
		if errs.Error != nil {
			logs.Logging.Error(ctx, errs.Error)
			goutil.ResponseError(c, http.StatusForbidden, errors.New(http.StatusText(http.StatusForbidden)), nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
