package database

import (
	"backend/models"
)

func (db *database) SelectServiceByName(serviceName string) (res models.Service, err error) {
	err = db.dbGorm.Where("name = ?", serviceName).First(&res).Error
	return
}

func (db *database) SelectAPIByName(name string) (res models.API, err error) {
	err = db.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (db *database) SelectAPIByEndpoint(endpoint string) (res models.API, err error) {
	err = db.dbGorm.Where(`endpoint = ?`, endpoint).First(&res).Error
	return
}
