package database

import (
	"backend/handler/security"
	"backend/handler/security/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type versionRepo struct {
	dbGorm *gorm.DB
}

func NewVersionRepo(dbGorm *gorm.DB) security.IVersionRepo {
	return &versionRepo{dbGorm: dbGorm}
}

func (v *versionRepo) InsertVersion(req models.Version) (res models.Version, err error) {
	err = v.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (v *versionRepo) SelectVersionByVersion(version string) (res models.Version, err error) {
	err = v.dbGorm.Where("version = ?", version).First(&res).Error
	return
}

func (v *versionRepo) UpdateVersion(req models.Version) (res models.Version, err error) {
	err = v.dbGorm.Model(&models.Version{}).Where("id = ?", req.ID).Updates(req).First(&res).Error
	return
}
