package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/tracer"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type versionRepo struct {
	dbGorm *gorm.DB
}

func NewVersionRepo(dbGorm *gorm.DB) security.IVersionRepo {
	return &versionRepo{dbGorm: dbGorm}
}

func (v *versionRepo) InsertVersion(ctx context.Context, req models.Version) (res models.Version, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: InsertVersion")
	defer span.End()
	err = v.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (v *versionRepo) SelectVersionByVersion(ctx context.Context, version string) (res models.Version, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectVersionByVersion")
	defer span.End()
	err = v.dbGorm.Where("version = ?", version).First(&res).Error
	return
}

func (v *versionRepo) UpdateVersion(ctx context.Context, req models.Version) (res models.Version, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: UpdateVersion")
	defer span.End()
	err = v.dbGorm.Model(&models.Version{}).Where("id = ?", req.ID).Updates(req).First(&res).Error
	return
}
