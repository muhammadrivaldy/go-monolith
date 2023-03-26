package entity

import (
	"backend/config"
	"backend/handler/template"
	"backend/handler/template/entity/database"
	"backend/logs"
	"context"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

type TemplateEntity struct {
	TemplateRepo template.ITemplateRepo
}

func NewTemplateEntity(conf config.Configuration) (TemplateEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Schema.Security.Address,
		conf.Database.Schema.Security.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return TemplateEntity{}, err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return TemplateEntity{}, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		conf.Database.Schema.Security.MigrationPath,
		conf.Database.Schema.Security.Database,
		driver)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return TemplateEntity{}, err
	}

	// do a migration up
	m.Up()

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return TemplateEntity{}, err
	}

	return TemplateEntity{
		TemplateRepo: database.NewTemplateRepo(dbGorm),
	}, nil
}
