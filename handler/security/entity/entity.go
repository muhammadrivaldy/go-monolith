package entity

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/entity/database"
	"backend/logs"
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

type SecurityEntity struct {
	AccessRepo  security.IAccessRepo
	ApiRepo     security.IApiRepo
	ServiceRepo security.IServiceRepo
}

func NewEntity(conf config.Configuration) (SecurityEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Schema.Security.Address,
		conf.Database.Schema.Security.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return SecurityEntity{}, err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return SecurityEntity{}, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		conf.Database.Schema.Security.MigrationPath,
		conf.Database.Schema.Security.Database,
		driver)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return SecurityEntity{}, err
	}

	// do a migration up
	m.Up()

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return SecurityEntity{}, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Address,
		DB:   0,
	})

	return SecurityEntity{
		AccessRepo:  database.NewAccessRepo(dbGorm, redisClient),
		ApiRepo:     database.NewApiRepo(dbGorm),
		ServiceRepo: database.NewServiceRepo(dbGorm),
	}, nil
}
