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

func (e endpoint) GetAccessApi(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: GetAccessApi")
	defer span.End()

	// get payload
	payload := payload.RequestGetAccessApi{UserType: util.StringToInt(c.Param("user_type"))}

	// validate
	if validationErrors := e.validation.ValidationStruct(payload); validationErrors.IsErrorExists() {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, validationErrors)
		return
	}

	// call service
	res, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return e.useCaseSecurity.GetAccessApi(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, res)

}
