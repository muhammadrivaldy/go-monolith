package rest

import (
	"backend/middleware"
	"backend/tracer"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e endpoint) GetServices(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: GetServices")
	defer span.End()

	// call services
	res, errs := middleware.WrapUseCase(ctx, nil, func() (interface{}, util.Error) {
		return e.useCaseSecurity.GetServices(ctx)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, res)

}
