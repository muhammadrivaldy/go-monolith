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
	usec       security.Usecase
	uhel       health.Usecase
	validation goutil.Validation
}

// NewEndpoint is a function for override a usecase method
func NewEndpoint(
	usec security.Usecase,
	uhel health.Usecase,
	engine *gin.Engine,
	validation goutil.Validation) error {

	// declare variable
	var edp = endpoint{
		usec:       usec,
		uhel:       uhel,
		validation: validation}

	// register the service
	serviceID, err := usec.RegisterService(context.TODO(), "Security")
	if err.Error != nil {
		return err.Error
	}

	// declare the endpoint
	const rootEndpoint = "/api/v1/health"
	var health = payload.RegisterAPIRequest{Name: "Health", Endpoint: rootEndpoint, Method: http.MethodGet, ServiceID: serviceID}

	// append data apis
	var APIs []*payload.RegisterAPIRequest
	APIs = append(APIs, &health)

	// register the apis
	for _, i := range APIs {
		usec.RegisterAPI(i)
	}

	// route the endpoint
	engine.Handle(health.Method, health.Endpoint, edp.Health)

	// send result
	return nil
}
