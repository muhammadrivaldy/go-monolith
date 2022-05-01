package database

import (
	"backend/config"
	"backend/handler/security"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

type database struct {
	dbGorm *gorm.DB
}

func NewDatabase(conf config.Configuration) (security.Database, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Schema.Security,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./handler/security/entity/database/migrations",
		conf.Database.Schema.Security,
		driver)
	if err != nil {
		return nil, err
	}

	m.Up()

	clientGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		return nil, err
	}

	return &database{dbGorm: clientGorm}, nil

}
