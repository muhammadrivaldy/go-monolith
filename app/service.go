package main

import (
	"backend/config"
	securityRest "backend/handler/security/delivery/rest"
	securityEntity "backend/handler/security/entity"
	securityUseCase "backend/handler/security/usecase"
	usersEntity "backend/handler/users/entity"
	usersUseCase "backend/handler/users/usecase"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func service(
	route *gin.Engine,
	config config.Configuration,
	validate goutil.Validation) {

	// call the function of method entity
	securityEntity, err := securityEntity.NewEntity(config)
	if err != nil {
		panic(err)
	}

	usersEntity, err := usersEntity.NewEntity(config)
	if err != nil {
		panic(err)
	}

	// call the function of method usecase
	securityUseCase := securityUseCase.NewUseCase(config, securityEntity, usersEntity)
	usersUseCase := usersUseCase.NewUseCase(config, usersEntity)

	// call the function of method endpoint
	securityRest.NewEndpoint(config, route, securityUseCase, usersUseCase, validate)
}
