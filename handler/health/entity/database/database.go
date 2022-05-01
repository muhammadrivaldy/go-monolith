package database

import (
	"backend/config"
	health "backend/handler/health"
	"strings"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

type database struct {
	dbGorm *gorm.DB
}

func NewDatabase(conf config.Configuration) (health.Database, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Schema.Security,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		return nil, err
	}

	clientGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		return nil, err
	}

	return &database{dbGorm: clientGorm}, nil

}

func (d *database) SelectTime() (res time.Time, err error) {
	err = d.dbGorm.Raw(`select now()`).Scan(&res).Error
	return
}
