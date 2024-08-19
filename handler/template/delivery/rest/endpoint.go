package rest

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/payload"
	"backend/handler/template"
	"backend/handler/users"
	"backend/middleware"
	"backend/util"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type endpoint struct {
	validation      goutil.Validation
	useCaseTemplate template.ITemplateUseCase
}

// NewEndpoint is a function for override a usecase method
func NewEndpoint(
	config config.Configuration,
	useCaseUsers users.IUserUseCase,
	useCaseSecurity security.ISecurityUseCase,
	useCaseTemplate template.ITemplateUseCase,
	engine *gin.Engine,
	validation goutil.Validation) error {

	// declare variable
	var _ = endpoint{
		validation:      validation,
		useCaseTemplate: useCaseTemplate,
	}

	// register the service
	serviceID, err := useCaseSecurity.RegisterService(context.TODO(), "Orders")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/templates"
	var createTemplate = payload.RequestRegisterApi{ID: "", Name: "Create Template", Endpoint: rootEndpoint + "/", Method: http.MethodPost, ServiceID: serviceID}

	// append data apis
	var APIs []*payload.RequestRegisterApi
	APIs = append(APIs, &createTemplate)

	// register the apis
	for _, i := range APIs {
		useCaseSecurity.RegisterApi(context.TODO(), i)
	}

	// middleware validate access
	middleware := middleware.NewMiddleware(useCaseUsers, useCaseSecurity)

	// route the endpoint
	engine.Use(util.EnableCORS)
	engine.Handle(createTemplate.Method, createTemplate.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(createTemplate.ID), func(ctx *gin.Context) {})

	// send result
	return nil
}
