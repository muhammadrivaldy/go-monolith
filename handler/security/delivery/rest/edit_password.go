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

func (e *endpoint) EditPassword(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: EditPassword")
	defer span.End()

	// set payload
	payload := payload.RequestEditPassword{
		UserID:   util.StringToInt64(c.Param("user_id")),
		Password: c.PostForm("password"),
	}

	// validate payload
	if validationErrors := e.validation.ValidationStruct(payload); validationErrors.IsErrorExists() {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, validationErrors)
		return
	}

	// call service
	response, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		errs := e.useCaseSecurity.EditPassword(ctx, payload)
		return nil, errs
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, response)
}
