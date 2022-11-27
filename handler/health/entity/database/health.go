package database

import (
	health "backend/handler/health"
	"time"

	"gorm.io/gorm"
)

type database struct {
	dbGorm *gorm.DB
}

func NewHealthRepo(dbGorm *gorm.DB) health.IHealthRepo {
	return &database{dbGorm: dbGorm}
}

func (d *database) SelectTime() (res time.Time, err error) {
	err = d.dbGorm.Raw(`select now()`).Scan(&res).Error
	return
}
