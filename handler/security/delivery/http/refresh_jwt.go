package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e *endpoint) RefreshJWT(c *gin.Context) {

	response, errs := e.useCaseSecurity.RefreshJWT(goutil.ParseContext(c))
	if errs.Error != nil {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, response)
}
