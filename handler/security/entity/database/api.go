package database

import (
	"backend/handler/security"
	"backend/handler/security/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type api struct {
	dbGorm *gorm.DB
}

func NewApiRepo(dbGorm *gorm.DB) security.IApiRepo {
	return api{dbGorm: dbGorm}
}

func (a api) InsertApi(req models.Api) (res models.Api, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a api) SelectApiByName(name string) (res models.Api, err error) {
	err = a.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (a api) SelectApiByEndpoint(endpoint string) (res models.Api, err error) {
	err = a.dbGorm.Where("endpoint = ?", endpoint).First(&res).Error
	return
}

func (a api) UpdateApi(req models.Api) (res models.Api, err error) {
	err = a.dbGorm.Model(&models.Api{}).Where(`id = ?`, req.ID).Updates(req).First(&res).Error
	return
}
