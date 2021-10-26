package security

import (
	"backend/models"
	"context"
	"sync"
)

// Usecase is a interface for layer business
type Usecase interface {
	RegisterAPI(wg *sync.WaitGroup, req *models.RegisterAPIRequest)
	RegisterService(ctx context.Context, serviceName string) (id int, errs models.Error)
}

type Entity struct {
	Database
}

type Database interface {
	Reader
	Writer
}

// Reader is a interface for layer data reader
type Reader interface {
	SelectServiceByName(serviceName string) (res models.Service, err error)
	SelectAPIByName(name string) (res models.API, err error)
	SelectAPIByEndpoint(endpoint string) (res models.API, err error)
}

// Writer is a interface for layer data writer
type Writer interface {
	InsertService(req models.Service) (res models.Service, err error)
	InsertAPI(req models.API) (res models.API, err error)

	UpdateAPI(req models.API) (res models.API, err error)
}
