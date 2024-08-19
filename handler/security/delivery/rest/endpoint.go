package rest

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/payload"
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
	serviceID, err := useCaseSecurity.RegisterService(context.TODO(), "Security")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/security"
	var login = payload.RequestRegisterApi{ID: "9e9de361-a468-4251-a25b-23995477695f", Name: "Login", Endpoint: rootEndpoint + "/login", Method: http.MethodPost, ServiceID: serviceID}
	var refreshJwt = payload.RequestRegisterApi{ID: "71597833-0501-4ff1-9252-3cca6bacaa20", Name: "Refresh JWT", Endpoint: rootEndpoint + "/refresh", Method: http.MethodPost, ServiceID: serviceID}
	var getAccessApi = payload.RequestRegisterApi{ID: "3227b75a-1c52-401d-95b3-15ab62301b4b", Name: "Get Access Api", Endpoint: rootEndpoint + "/accesses/users/:user_type", Method: http.MethodGet, ServiceID: serviceID}
	var patchAccessApi = payload.RequestRegisterApi{ID: "46bfa26d-0b7f-401e-91b0-060c8631efc7", Name: "Patch Access Api", Endpoint: rootEndpoint + "/accesses/users/:user_type", Method: http.MethodPatch, ServiceID: serviceID}
	var getServices = payload.RequestRegisterApi{ID: "db29f615-432d-4044-943e-95b0c96f69f7", Name: "Get Services", Endpoint: rootEndpoint + "/services", Method: http.MethodGet, ServiceID: serviceID}
	var getApisByServiceID = payload.RequestRegisterApi{ID: "ab246918-33f0-4881-a8ee-7d6f198e6b94", Name: "Get Apis By Service ID", Endpoint: rootEndpoint + "/services/:service_id/api", Method: http.MethodGet, ServiceID: serviceID}
	var versionSupport = payload.RequestRegisterApi{ID: "f30bdce4-22eb-4387-9315-7b6c7f52e63d", Name: "Version Support", Endpoint: rootEndpoint + "/version/support", Method: http.MethodPost, ServiceID: serviceID}
	var updatePassword = payload.RequestRegisterApi{ID: "93c47cfa-9a68-4a74-88d8-4bcd32ab7efe", Name: "Update Password", Endpoint: rootEndpoint + "/users/:user_id/password", Method: http.MethodPut, ServiceID: serviceID}

	// append data apis
	var Apis []*payload.RequestRegisterApi
	Apis = append(Apis,
		&login,
		&refreshJwt,
		&getAccessApi,
		&patchAccessApi,
		&getServices,
		&getApisByServiceID,
		&versionSupport,
		&updatePassword)

	// register the apis
	for _, i := range Apis {
		go useCaseSecurity.RegisterApi(context.TODO(), i)
	}

	// middleware validate access
	middleware := middleware.NewMiddleware(useCaseUser, useCaseSecurity)

	// route the endpoint
	engine.Handle(login.Method, login.Endpoint, edp.Login)
	engine.Handle(refreshJwt.Method, refreshJwt.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(refreshJwt.ID), edp.RefreshJWT)
	engine.Handle(getAccessApi.Method, getAccessApi.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(getAccessApi.ID), edp.GetAccessApi)
	engine.Handle(patchAccessApi.Method, patchAccessApi.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(patchAccessApi.ID), edp.PatchAccessApi)
	engine.Handle(getServices.Method, getServices.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(getServices.ID), edp.GetServices)
	engine.Handle(getApisByServiceID.Method, getApisByServiceID.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(getApisByServiceID.ID), edp.GetApisByServiceID)
	engine.Handle(versionSupport.Method, versionSupport.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(versionSupport.ID), edp.VersionSupport)
	engine.Handle(updatePassword.Method, updatePassword.Endpoint, goutil.ParseJWT(config.JWTKey, jwt.SigningMethodHS256, util.AttributesJWT), middleware.ValidateAccess(updatePassword.ID), edp.EditPassword)

	// send result
	return nil
}
