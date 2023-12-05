package rest

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
	var login = payload.RequestRegisterApi{Name: "Login", Endpoint: rootEndpoint + "/login", Method: http.MethodPost, ServiceID: serviceId}
	var refreshJwt = payload.RequestRegisterApi{Name: "Refresh JWT", Endpoint: rootEndpoint + "/refresh", Method: http.MethodPost, ServiceID: serviceId}
	var getAccessApi = payload.RequestRegisterApi{Name: "Get Access Api", Endpoint: rootEndpoint + "/accesses/users/:user_type", Method: http.MethodGet, ServiceID: serviceId}
	var patchAccessApi = payload.RequestRegisterApi{Name: "Patch Access Api", Endpoint: rootEndpoint + "/accesses/users/:user_type", Method: http.MethodPatch, ServiceID: serviceId}
	var getServices = payload.RequestRegisterApi{Name: "Get Services", Endpoint: rootEndpoint + "/services", Method: http.MethodGet, ServiceID: serviceId}
	var getApisByServiceId = payload.RequestRegisterApi{Name: "Get Apis By Service Id", Endpoint: rootEndpoint + "/services/:service_id/api", Method: http.MethodGet, ServiceID: serviceId}
	var versionSupport = payload.RequestRegisterApi{Name: "Version Support", Endpoint: rootEndpoint + "/version/support", Method: http.MethodPost, ServiceID: serviceId}
	var updatePassword = payload.RequestRegisterApi{Name: "Update Password", Endpoint: rootEndpoint + "/users/:user_id/password", Method: http.MethodPut, ServiceID: serviceId}

	// append data apis
	var Apis []*payload.RequestRegisterApi
	Apis = append(Apis,
		&login,
		&refreshJwt,
		&getAccessApi,
		&patchAccessApi,
		&getServices,
		&getApisByServiceId,
		&versionSupport,
		&updatePassword)

	// register the apis
	for _, i := range Apis {
		useCaseSecurity.RegisterApi(context.TODO(), i)
	}

	// middleware validate access
	middleware := middleware.NewMiddleware(useCaseUser, useCaseSecurity)

	// route the endpoint
	engine.Handle(login.Method, login.Endpoint, edp.Login)
	engine.Handle(refreshJwt.Method, refreshJwt.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(refreshJwt.ID), edp.RefreshJWT)
	engine.Handle(getAccessApi.Method, getAccessApi.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(getAccessApi.ID), edp.GetAccessApi)
	engine.Handle(patchAccessApi.Method, patchAccessApi.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(patchAccessApi.ID), edp.PatchAccessApi)
	engine.Handle(getServices.Method, getServices.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(getServices.ID), edp.GetServices)
	engine.Handle(getApisByServiceId.Method, getApisByServiceId.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(getApisByServiceId.ID), edp.GetApisByServiceId)
	engine.Handle(versionSupport.Method, versionSupport.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(versionSupport.ID), edp.VersionSupport)
	engine.Handle(updatePassword.Method, updatePassword.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256), middleware.ValidateAccess(updatePassword.ID), edp.UpdatePassword)

	// send result
	return nil
}
