package main

import (
	"backend/config"
	healthHttp "backend/handler/health/delivery/http"
	healthEnti "backend/handler/health/entity"
	healthUc "backend/handler/health/usecase"
	securityEnti "backend/handler/security/entity"
	securityUc "backend/handler/security/usecase"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func service(
	route *gin.Engine,
	config config.Configuration,
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

	// call the function of method useCase
	healthuseCase := healthUc.NewuseCase(logs, healthEntity)
	securityuseCase := securityUc.NewuseCase(logs, config, securityEntity)

	// call the function of method endpoint
	healthHttp.NewEndpoint(securityuseCase, healthuseCase, route, validate)
}
