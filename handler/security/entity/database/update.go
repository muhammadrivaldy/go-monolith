package database

import "backend/handler/security/models"

func (db *database) UpdateAPI(req models.API) (res models.API, err error) {
	err = db.dbGorm.Model(&models.API{}).Where(`id = ?`, req.ID).Updates(req).First(&res).Error
	return
}
