package security

import (
	"backend/handler/security/models"
	"backend/handler/security/payload"
	"backend/util"
	"context"
)

// ISecurityUseCase is a interface for layer business
type ISecurityUseCase interface {
	RegisterApi(ctx context.Context, req *payload.RequestRegisterApi)
	RegisterService(ctx context.Context, serviceName string) (id int64, errs util.Error)
	ValidateAccessUser(ctx context.Context, apiId string) (res bool, errs util.Error)
	Login(ctx context.Context, req payload.RequestLogin) (res payload.ResponseLogin, errs util.Error)
	RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error)
	GetAccessApi(ctx context.Context, req payload.RequestGetAccessApi) (res payload.ResponseGetAccessApi, errs util.Error)
	PatchAccessApi(ctx context.Context, req payload.RequestPatchAccessApi) (errs util.Error)
	GetServices(ctx context.Context) (res []payload.ResponseGetServices, errs util.Error)
	GetApisByServiceId(ctx context.Context, req payload.RequestGetApisServiceId) (res []payload.ResponseGetApisServiceId, errs util.Error)
	VersionSupport(ctx context.Context, req payload.RequestVersionSupport) (res payload.ResponseVersionSupport, errs util.Error)
	EditPassword(ctx context.Context, req payload.RequestEditPassword) (errs util.Error)
}

type IApiRepo interface {
	InsertApi(req models.Api) (res models.Api, err error)
	SelectApiByID(ctx context.Context, id string) (res models.Api, err error)
	SelectApiByName(name string) (res models.Api, err error)
	SelectApiByEndpoint(endpoint, method string) (res models.Api, err error)
	SelectApisByServiceId(serviceId int) (res []models.Api, err error)
	UpdateApi(req models.Api) (res models.Api, err error)
}

type IServiceRepo interface {
	InsertService(req models.Service) (res models.Service, err error)
	SelectServiceById(id int) (res models.Service, err error)
	SelectServiceByName(name string) (res models.Service, err error)
	SelectServices() (res []models.Service, err error)
}

type IAccessRepo interface {
	InsertAccess(req models.Access) (res models.Access, err error)
	InsertAccesses(req []models.Access) (res []models.Access, err error)
	SelectAccessByFilter(ctx context.Context, req util.FilterQuery) (res models.Access, err error)
	SelectAccessByUserType(userTypeId int) (res []models.Access, err error)
	UpdateAccess(req models.Access) (res models.Access, err error)
	DeleteAccessesByUserTypeIdAndApiId(userTypeId int, apiId []int) (err error)
}

type IVersionRepo interface {
	InsertVersion(req models.Version) (res models.Version, err error)
	SelectVersionByVersion(version string) (res models.Version, err error)
	UpdateVersion(req models.Version) (res models.Version, err error)
}
