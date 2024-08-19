package middleware

import (
	"backend/handler/security"
	"backend/handler/users"
	"backend/handler/users/payload"
	"backend/util"
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

func (m middleware) ValidateAccess(apiID string) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := goutil.ParseContext(c)

		userInfo := util.GetContext(ctx)

		modelUser, errs := m.useCaseUser.GetUserById(ctx, payload.RequestGetUserById{UserId: userInfo.UserId})
		if errs.Error != nil {
			goutil.ResponseError(c, errs.Code, errs.Error, errs.Object)
			c.Abort()
			return
		}

		// the UUID (602587ea-2364-4000-9d5e-4ca0c4bf8d7c) is id of refresh jwt endpoint
		// if the apiID is similar with that, we need to validate the token-type and ensure the
		// token is token for refresh new token. We can check it based on the token type
		if apiID == "602587ea-2364-4000-9d5e-4ca0c4bf8d7c" {
			if userInfo.Type != goutil.RefreshToken {
				goutil.ResponseError(c, http.StatusUnauthorized, util.ErrorUnauthorized, nil)
				c.Abort()
				return
			}
		} else {
			if int(modelUser.UserType) != userInfo.UserType {
				goutil.ResponseError(c, http.StatusUnauthorized, util.ErrorUnauthorized, nil)
				c.Abort()
				return
			}
		}

		_, errs = m.useCaseSecurity.ValidateAccessUser(ctx, apiID)
		if errs.Error != nil {
			goutil.ResponseError(c, http.StatusForbidden, errors.New("user does not allowed to access this endpoint"), nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
