package http

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e endpoint) GetAccessApi(c *gin.Context) {

	// get payload
	payload := payload.RequestGetAccessApi{UserType: util.StringToInt(c.Param("user_type"))}

	// validate
	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, nil)
		return
	}

	ctx := goutil.ParseContext(c)

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
