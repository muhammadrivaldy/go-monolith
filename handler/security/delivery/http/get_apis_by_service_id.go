package http

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e endpoint) GetApisByServiceId(c *gin.Context) {

	// get payload
	payload := payload.RequestGetApisServiceId{
		ServiceID: util.StringToInt(c.Param("service_id")),
	}

	ctx := goutil.ParseContext(c)

	// call service
	res, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return e.useCaseSecurity.GetApisByServiceId(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, res)

}
