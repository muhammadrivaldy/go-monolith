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
	ValidateAccessUser(ctx context.Context, apiID string) (res bool, errs util.Error)
	Login(ctx context.Context, req payload.RequestLogin) (res payload.ResponseLogin, errs util.Error)
	RefreshJWT(ctx context.Context) (res payload.ResponseLogin, errs util.Error)
	GetAccessApi(ctx context.Context, req payload.RequestGetAccessApi) (res payload.ResponseGetAccessApi, errs util.Error)
	PatchAccessApi(ctx context.Context, req payload.RequestPatchAccessApi) (errs util.Error)
	GetServices(ctx context.Context) (res []payload.ResponseGetServices, errs util.Error)
	GetApisByServiceID(ctx context.Context, req payload.RequestGetApisServiceID) (res []payload.ResponseGetApisServiceID, errs util.Error)
	VersionSupport(ctx context.Context, req payload.RequestVersionSupport) (res payload.ResponseVersionSupport, errs util.Error)
	EditPassword(ctx context.Context, req payload.RequestEditPassword) (errs util.Error)
}

type IApiRepo interface {
	InsertApi(ctx context.Context, req models.Api) (res models.Api, err error)
	SelectApiByID(ctx context.Context, id string) (res models.Api, err error)
	SelectApiByName(ctx context.Context, name string) (res models.Api, err error)
	SelectApiByEndpoint(ctx context.Context, endpoint, method string) (res models.Api, err error)
	SelectApisByServiceID(ctx context.Context, serviceID int) (res []models.Api, err error)
	UpdateApi(ctx context.Context, req models.Api) (res models.Api, err error)
}

type IServiceRepo interface {
	InsertService(ctx context.Context, req models.Service) (res models.Service, err error)
	SelectServiceByID(ctx context.Context, id int) (res models.Service, err error)
	SelectServiceByName(ctx context.Context, name string) (res models.Service, err error)
	SelectServices(ctx context.Context) (res []models.Service, err error)
}

type IAccessRepo interface {
	InsertAccess(ctx context.Context, req models.Access) (res models.Access, err error)
	InsertAccesses(ctx context.Context, req []models.Access) (res []models.Access, err error)
	SelectAccessByFilter(ctx context.Context, req util.FilterQuery) (res models.Access, err error)
	SelectAccessByUserType(ctx context.Context, userTypeID int) (res []models.Access, err error)
	UpdateAccess(ctx context.Context, req models.Access) (res models.Access, err error)
	DeleteAccessesByUserTypeIDAndApiID(ctx context.Context, userTypeID int, apiID []int) (err error)
}

type IVersionRepo interface {
	InsertVersion(ctx context.Context, req models.Version) (res models.Version, err error)
	SelectVersionByVersion(ctx context.Context, version string) (res models.Version, err error)
	UpdateVersion(ctx context.Context, req models.Version) (res models.Version, err error)
}
