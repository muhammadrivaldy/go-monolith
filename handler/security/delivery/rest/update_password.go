package rest

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) UpdatePassword(c *gin.Context) {

	// set payload
	payload := payload.RequestUpdatePassword{
		UserId:   util.StringToInt(c.Param("user_id")),
		Password: c.PostForm("password"),
	}

	// validate payload
	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, nil)
		return
	}

	ctx := goutil.ParseContext(c)

	// call service
	response, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		errs := e.useCaseSecurity.UpdatePassword(ctx, payload)
		return nil, errs
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, response)
}
