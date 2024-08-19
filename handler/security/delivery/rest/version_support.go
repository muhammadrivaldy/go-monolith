package rest

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/tracer"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) VersionSupport(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: VersionSupport")
	defer span.End()

	// set payload
	payload := payload.RequestVersionSupport{
		Version: c.PostForm("version"),
	}

	// validate payload
	if validationErrors := e.validation.ValidationStruct(payload); validationErrors.IsErrorExists() {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, validationErrors)
		return
	}

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
