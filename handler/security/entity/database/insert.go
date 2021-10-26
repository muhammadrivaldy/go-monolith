package database

import (
	"backend/models"

	"gorm.io/gorm/clause"
)

func (db *database) InsertService(req models.Service) (res models.Service, err error) {
	err = db.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (db *database) InsertAPI(req models.API) (res models.API, err error) {
	err = db.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}
