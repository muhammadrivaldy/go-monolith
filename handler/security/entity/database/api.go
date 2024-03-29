package database

import (
	"backend/handler/security"
	"backend/handler/security/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type apiRepo struct {
	dbGorm *gorm.DB
}

func NewApiRepo(dbGorm *gorm.DB) security.IApiRepo {
	return apiRepo{dbGorm: dbGorm}
}

func (a apiRepo) InsertApi(req models.Api) (res models.Api, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a apiRepo) SelectApiByName(name string) (res models.Api, err error) {
	err = a.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (a apiRepo) SelectApiByEndpoint(endpoint, method string) (res models.Api, err error) {
	err = a.dbGorm.Where("endpoint = ? and method = ?", endpoint, method).First(&res).Error
	return
}

func (a apiRepo) SelectApisByServiceId(serviceId int) (res []models.Api, err error) {
	err = a.dbGorm.Where("service_id = ?", serviceId).Find(&res).Error
	return
}

func (a apiRepo) UpdateApi(req models.Api) (res models.Api, err error) {
	err = a.dbGorm.Model(&models.Api{}).Where("id = ?", req.Id).Updates(req).First(&res).Error
	return
}
