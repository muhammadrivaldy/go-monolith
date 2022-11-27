package database

import (
	"backend/handler/security"
	"backend/handler/security/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type service struct {
	dbGorm *gorm.DB
}

func NewServiceRepo(dbGorm *gorm.DB) security.IServiceRepo {
	return service{dbGorm: dbGorm}
}

func (s service) InsertService(req models.Service) (res models.Service, err error) {
	err = s.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (s service) SelectServiceByName(name string) (res models.Service, err error) {
	err = s.dbGorm.Where("name = ?", name).First(&res).Error
	return
}
