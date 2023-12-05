package rest

import (
	"backend/handler/security/payload"
	"backend/middleware"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) Login(c *gin.Context) {

	email, password, _ := c.Request.BasicAuth()

	// set payload
	payload := payload.RequestLogin{
		Email:    email,
		Password: password,
	}

	// validate payload
	if err := e.validation.ValidationStruct(payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, util.ErrorIncorrectInput, nil)
		return
	}

	ctx := goutil.ParseContext(c)

	// call service
	response, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return e.useCaseSecurity.Login(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, response)
}
