package security

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/util"
	"context"
)

// UseCase is a interface for layer business
type ISecurityUseCase interface {
	RegisterApi(req *payload.RegisterApiRequest)
	RegisterService(ctx context.Context, serviceName string) (id int, errs util.Error)
}

type IApiRepo interface {
	InsertApi(req models.Api) (res models.Api, err error)
	SelectApiByName(name string) (res models.Api, err error)
	SelectApiByEndpoint(endpoint string) (res models.Api, err error)
	UpdateApi(req models.Api) (res models.Api, err error)
}

type IServiceRepo interface {
	InsertService(req models.Service) (res models.Service, err error)
	SelectServiceByName(name string) (res models.Service, err error)
}

type IAccessRepo interface {
	InsertAccess(req models.Access) (res models.Access, err error)
	SelectAccessByFilter(req util.FilterQuery) (res models.Access, err error)
	UpdateAccess(req models.Access) (res models.Access, err error)
}
