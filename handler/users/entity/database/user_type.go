package database

import (
	"backend/handler/users"
	"backend/handler/users/models"

	"gorm.io/gorm"
)

type userTypeRepo struct {
	dbGorm *gorm.DB
}

func NewUserTypeRepo(dbGorm *gorm.DB) users.IUserTypeRepo {
	return userTypeRepo{dbGorm: dbGorm}
}

func (u userTypeRepo) SelectUserTypeByID(id int) (res models.UserType, err error) {
	err = u.dbGorm.Where("id = ?", id).First(&res).Error
	return
}
