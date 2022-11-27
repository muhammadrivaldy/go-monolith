package http

import (
	health "backend/handler/health"
	"backend/handler/security"
	"backend/handler/security/payload"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

type endpoint struct {
	useCaseSecurity security.ISecurityUseCase
	useCaseHealth   health.IHealthUseCase
	validation      goutil.Validation
}

// NewEndpoint is a function for override a useCase method
func NewEndpoint(
	engine *gin.Engine,
	useCaseSecurity security.ISecurityUseCase,
	useCaseHealth health.IHealthUseCase,
	validation goutil.Validation) error {

	// declare variable
	var edp = endpoint{
		useCaseSecurity: useCaseSecurity,
		useCaseHealth:   useCaseHealth,
		validation:      validation}

	// register the service
	serviceID, err := useCaseSecurity.RegisterService(context.TODO(), "Security")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/health"
	var health = payload.RegisterApiRequest{Name: "Health", Endpoint: rootEndpoint, Method: http.MethodGet, ServiceID: serviceID}

	// append data apis
	var Apis []*payload.RegisterApiRequest
	Apis = append(Apis, &health)

	// register the apis
	for _, i := range Apis {
		useCaseSecurity.RegisterApi(i)
	}

	// route the endpoint
	engine.Handle(health.Method, health.Endpoint, edp.Health)

	// send result
	return nil
}
