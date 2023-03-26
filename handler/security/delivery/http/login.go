package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) Login(c *gin.Context) {

	email, password, _ := c.Request.BasicAuth()

	if err := e.validation.ValidationVariable(email, "email", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	if err := e.validation.ValidationVariable(email, "required,min=8", ""); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)), nil)
		return
	}

	response, errs := e.useCaseSecurity.Login(context.Background(), email, password)
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, response)
}
