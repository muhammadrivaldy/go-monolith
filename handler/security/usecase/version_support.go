package usecase

import (
	"backend/handler/security/payload"
	"backend/logs"
	"backend/util"
	"context"

	"gorm.io/gorm"
)

func (s *securityUseCase) VersionSupport(ctx context.Context, req payload.RequestVersionSupport) (res payload.ResponseVersionSupport, errs util.Error) {

	// get version
	version, err := s.securityEntity.VersionRepo.SelectVersionByVersion(req.Version)
	if err == gorm.ErrRecordNotFound {
		return payload.ResponseVersionSupport{
			Version: req.Version,
			Support: false,
		}, util.ErrorMapping(nil)
	} else if err != nil {
		logs.Logging.Error(ctx, err)
		return payload.ResponseVersionSupport{Version: version.Version, Support: version.Support}, util.ErrorMapping(err)
	}

	return payload.ResponseVersionSupport{
			Version: version.Version,
			Support: version.Support},
		util.ErrorMapping(nil)

}
