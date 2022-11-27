package main

import (
	"backend/config"
	healthHttp "backend/handler/health/delivery/http"
	healthEntity "backend/handler/health/entity"
	healthUc "backend/handler/health/usecase"
	securityEntity "backend/handler/security/entity"
	securityUc "backend/handler/security/usecase"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func service(
	route *gin.Engine,
	config config.Configuration,
	validate goutil.Validation) {

	// call the function of method entity
	securityEntity, err := securityEntity.NewSecurityEntity(config)
	if err != nil {
		panic(err)
	}

	healthEntity, err := healthEntity.NewEntity(config)
	if err != nil {
		panic(err)
	}

	// call the function of method useCase
	healthUseCase := healthUc.NewUseCase(nil, healthEntity)
	securityUseCase := securityUc.NewSecurityUseCase(config, securityEntity)

	// call the function of method endpoint
	healthHttp.NewEndpoint(route, securityUseCase, healthUseCase, validate)
}
