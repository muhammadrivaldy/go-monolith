package main

import (
	healthHttp "backend/handler/health/delivery/http"
	healthEnti "backend/handler/health/entity"
	healthUc "backend/handler/health/usecase"
	securityEnti "backend/handler/security/entity"
	securityUc "backend/handler/security/usecase"
	"backend/models"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func service(
	route *gin.Engine,
	config models.Configuration,
	logs goutil.Logs,
	validate goutil.Validation) {

	// call the function of method entity
	securityEntity, err := securityEnti.NewEntity(config)
	if err != nil {
		panic(err)
	}

	healthEntity, err := healthEnti.NewEntity(config)
	if err != nil {
		panic(err)
	}

	// call the function of method usecase
	healthUsecase := healthUc.NewUsecase(logs, healthEntity)
	securityUsecase := securityUc.NewUsecase(logs, config, securityEntity)

	// call the function of method endpoint
	healthHttp.NewEndpoint(securityUsecase, healthUsecase, route, validate)
}
