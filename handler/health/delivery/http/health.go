package http

import (
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) Health(c *gin.Context) {

	_, errs := e.useCaseHealth.HealthService(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	_, errs = e.useCaseHealth.HealthDB(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, nil)
}
