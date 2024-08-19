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

func (e endpoint) GetApisByServiceID(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: GetApisByServiceID")
	defer span.End()

	// get payload
	payload := payload.RequestGetApisServiceID{
		ServiceID: util.StringToInt(c.Param("service_id")),
	}

	// call service
	res, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return e.useCaseSecurity.GetApisByServiceID(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, res)

}
