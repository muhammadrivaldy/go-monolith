package entity

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/entity/database"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

type SecurityEntity struct {
	AccessRepo  security.IAccessRepo
	ApiRepo     security.IApiRepo
	ServiceRepo security.IServiceRepo
}

func NewSecurityEntity(conf config.Configuration) (SecurityEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Schema.Security,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		return SecurityEntity{}, err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		return SecurityEntity{}, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./handler/security/entity/database/migrations",
		conf.Database.Schema.Security,
		driver)
	if err != nil {
		return SecurityEntity{}, err
	}

	m.Up()

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		return SecurityEntity{}, err
	}

	return SecurityEntity{
		AccessRepo:  database.NewAccessRepo(dbGorm),
		ApiRepo:     database.NewApiRepo(dbGorm),
		ServiceRepo: database.NewServiceRepo(dbGorm),
	}, nil
}
