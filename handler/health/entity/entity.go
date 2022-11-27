package database

import (
	"backend/config"
	health "backend/handler/health"
	db "backend/handler/health/entity/database"
	"strings"

	goutil "github.com/muhammadrivaldy/go-util"
)

type HealthEntity struct {
	HealthRepo health.IHealthRepo
}

func NewEntity(conf config.Configuration) (HealthEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Schema.Security,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		return HealthEntity{}, err
	}

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		return HealthEntity{}, err
	}

	return HealthEntity{
		HealthRepo: db.NewHealthRepo(dbGorm),
	}, nil
}
