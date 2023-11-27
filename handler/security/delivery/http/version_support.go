package http

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) VersionSupport(c *gin.Context) {

	// set payload
	payload := payload.RequestVersionSupport{
		Version: c.PostForm("version"),
	}

	// validate payload
	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, err, nil)
		return
	}

	ctx := goutil.ParseContext(c)

	// call service
	response, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return e.useCaseSecurity.VersionSupport(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, response)

}
