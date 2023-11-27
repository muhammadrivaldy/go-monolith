package database

import (
	"backend/handler/security"
	"backend/handler/security/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceRepo struct {
	dbGorm *gorm.DB
}

func NewServiceRepo(dbGorm *gorm.DB) security.IServiceRepo {
	return serviceRepo{dbGorm: dbGorm}
}

func (s serviceRepo) InsertService(req models.Service) (res models.Service, err error) {
	err = s.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (s serviceRepo) SelectServiceById(id int) (res models.Service, err error) {
	err = s.dbGorm.Where("id = ?", id).First(&res).Error
	return
}

func (s serviceRepo) SelectServiceByName(name string) (res models.Service, err error) {
	err = s.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (s serviceRepo) SelectServices() (res []models.Service, err error) {
	err = s.dbGorm.Find(&res).Error
	return
}
