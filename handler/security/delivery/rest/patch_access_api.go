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

func (e endpoint) PatchAccessApi(c *gin.Context) {

	ctx := goutil.ParseContext(c)
	ctx, span := tracer.Tracer.Start(ctx, "REST: PatchAccessApi")
	defer span.End()

	// bind payload
	payload := payload.RequestPatchAccessApi{}
	if err := c.Bind(&payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, err, nil)
		return
	}

	payload.UserType = util.StringToInt(c.Param("user_type"))

	// call service
	_, errs := middleware.WrapUseCase(ctx, payload, func() (interface{}, util.Error) {
		return nil, e.useCaseSecurity.PatchAccessApi(ctx, payload)
	})
	if errs.IsError() {
		goutil.ResponseError(c, errs.Code, errs.Error, nil)
		return
	}

	// response
	goutil.ResponseOK(c, http.StatusOK, nil)

}
