package rest

import (
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) RefreshJWT(c *gin.Context) {

	ctx := goutil.ParseContext(c)

	res, errs := middleware.WrapUseCase(ctx, nil, func() (interface{}, util.Error) {
		return e.useCaseSecurity.RefreshJWT(ctx)
	})
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, res)
}
