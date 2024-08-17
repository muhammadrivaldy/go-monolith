package main

import (
	"backend/config"
	"backend/logs"
	"context"
	"strings"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/gorm"
)

func databaseClient(conf config.Configuration) (*gorm.DB, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return nil, err
	}

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return nil, err
	}

	dbGorm.Use(otelgorm.NewPlugin())

	return dbGorm, nil
}

func databaseMigration(conf config.Configuration) error {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Address,
		conf.Database.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		conf.Database.MigrationPath,
		conf.Database.Database,
		driver)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return err
	}

	// do a migration up
	return m.Up()
}

func redisClient(conf config.Configuration) *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Address,
		DB:   0,
	})

	redisClient.AddHook(redisotel.NewTracingHook())

	return redisClient
}
