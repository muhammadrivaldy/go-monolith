package http

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/payload"
	"backend/handler/users"
	"backend/middleware"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type endpoint struct {
	useCaseSecurity security.ISecurityUseCase
	validation      goutil.Validation
}

// NewEndpoint is a function for override a useCase method
func NewEndpoint(
	config config.Configuration,
	engine *gin.Engine,
	useCaseSecurity security.ISecurityUseCase,
	useCaseUser users.IUserUseCase,
	validation goutil.Validation) error {

	// declare variable
	var edp = endpoint{
		useCaseSecurity: useCaseSecurity,
		validation:      validation}

	// register the service
	serviceId, err := useCaseSecurity.RegisterService(context.TODO(), "Security")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/security"
	var login = payload.RegisterApiRequest{Name: "Login", Endpoint: rootEndpoint + "/login", Method: http.MethodPost, ServiceId: serviceId}
	var refreshJwt = payload.RegisterApiRequest{Name: "Refresh JWT", Endpoint: rootEndpoint + "/refresh", Method: http.MethodPost, ServiceId: serviceId}

	// append data apis
	var Apis []*payload.RegisterApiRequest
	Apis = append(Apis, &login, &refreshJwt)

	// register the apis
	for _, i := range Apis {
		useCaseSecurity.RegisterApi(context.TODO(), i)
	}

	// middleware validate access
	middleware := middleware.NewMiddleware(useCaseUser, useCaseSecurity)

	// route the endpoint
	engine.Handle(login.Method, login.Endpoint, edp.Login)
	engine.Handle(refreshJwt.Method, refreshJwt.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(refreshJwt.Id), edp.RefreshJWT)

	// send result
	return nil
}
