package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/util"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type access struct {
	dbGorm *gorm.DB
}

func NewAccessRepo(dbGorm *gorm.DB) security.IAccessRepo {
	return access{dbGorm: dbGorm}
}

func (a access) InsertAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a access) SelectAccessByFilter(req util.FilterQuery) (res models.Access, err error) {
	query, arguments := req.BuildQuery()
	err = a.dbGorm.Where(query, arguments...).First(&res).Error
	return
}

func (a access) UpdateAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Model(&models.Access{}).Where(`id = ?`, req.ID).Updates(req).First(&res).Error
	return
}
