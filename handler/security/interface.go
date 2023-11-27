package security

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/util"
	"context"
)

// ISecurityUseCase is a interface for layer business
type ISecurityUseCase interface {
	RegisterApi(ctx context.Context, req *payload.RegisterApiRequest)
	RegisterService(ctx context.Context, serviceName string) (id int, errs util.Error)
	ValidateAccessUser(ctx context.Context, apiId int64) (res bool, errs util.Error)
	Login(ctx context.Context, email, password string) (res payload.ResponseLogin, errs util.Error)
	RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error)
}

type IApiRepo interface {
	InsertApi(req models.Api) (res models.Api, err error)
	SelectApiByName(name string) (res models.Api, err error)
	SelectApiByEndpoint(endpoint, method string) (res models.Api, err error)
	UpdateApi(req models.Api) (res models.Api, err error)
}

type IServiceRepo interface {
	InsertService(req models.Service) (res models.Service, err error)
	SelectServiceByName(name string) (res models.Service, err error)
}

type IAccessRepo interface {
	InsertAccess(req models.Access) (res models.Access, err error)
	InsertAccesses(req []models.Access) (res []models.Access, err error)
	SelectAccessByFilter(req util.FilterQuery) (res models.Access, err error)
	SelectAccessByUserType(userTypeId int) (res []models.Access, err error)
	UpdateAccess(req models.Access) (res models.Access, err error)
	DeleteAccessesByUserTypeIdAndApiId(userTypeId int, apiId []int) (err error)
}

type IVersionRepo interface {
	InsertVersion(req models.Version) (res models.Version, err error)
	SelectVersionByVersion(version string) (res models.Version, err error)
	UpdateVersion(req models.Version) (res models.Version, err error)
}
