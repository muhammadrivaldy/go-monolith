package database

import (
	"backend/handler/users"
	"backend/handler/users/models"
	"backend/tracer"
	"context"

	"gorm.io/gorm"
)

type userTypeRepo struct {
	dbGorm *gorm.DB
}

func NewUserTypeRepo(dbGorm *gorm.DB) users.IUserTypeRepo {
	return userTypeRepo{dbGorm: dbGorm}
}

func (u userTypeRepo) SelectUserTypeByID(ctx context.Context, id int) (res models.UserType, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectUserTypeByID")
	defer span.End()
	err = u.dbGorm.Where("id = ?", id).First(&res).Error
	return
}
