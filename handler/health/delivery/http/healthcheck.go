package http

import (
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) HealthService(c *gin.Context) {
	res, errs := e.uhel.HealthService(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, res)
}

func (e *endpoint) HealthDB(c *gin.Context) {
	res, errs := e.uhel.HealthDB(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, res)
}
