package main

import (
	"backend/config"
	"backend/logs"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
	"github.com/robfig/cron/v3"
)

const (
	pathConfigEnv = "PATH_CONF"
	fileConfigEnv = "FILE_CONF"
)

func main() {

	// get path & file configuration from environment variable
	pathConfig := os.Getenv(pathConfigEnv)
	fileConfig := os.Getenv(fileConfigEnv)

	// declare variable
	var err error

	// logging
	var osLog *os.File
	defer osLog.Close()
	osLog, err = goutil.OpenFile("./logs", filenameLog())
	if err != nil {
		panic(err)
	}

	// override io writer to gin default writer
	gin.DefaultWriter = io.Writer(osLog)

	// cron for create a log file every day
	c := cron.New()
	c.AddFunc("@midnight", func() {
		osLogNewest, err := goutil.OpenFileNewest(osLog, "./logs", filenameLog())
		if err != nil {
			return
		}

		logs.Logging.Config(osLogNewest)
		gin.DefaultWriter = io.Writer(osLogNewest)
		osLog = osLogNewest
	})
	c.Start()

	// open file for configuration
	osFile, err := goutil.OpenFile(pathConfig, fileConfig)
	if err != nil {
		panic(err)
	}
	defer osFile.Close()

	// get config
	var config config.Configuration
	if err := goutil.Configuration(osFile, &config); err != nil {
		panic(err)
	}

	// third-party telegram
	telegram, err := goutil.NewTele(config.ThirdParty.Telegram.Token, config.ThirdParty.Telegram.ChatID)
	if err != nil {
		panic(err)
	}

	// logs service
	logs.Logging = goutil.NewLog(osLog, telegram)
	defer logs.Logging.Sync()
	defer logs.Logging.Undo()

	logs.Logging.Info(context.Background(), "Success Load Configuration")

	// call gin route
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(requestid.New())
	route.Use(goutil.SetContext)
	route.Use(goutil.LogMiddleware(logs.Logging))

	// don't remove this code
	route.Static("nominatim", "../nominatim")

	// validation method
	validate, err := goutil.NewValidation()
	if err != nil {
		panic(err)
	}

	// running the service
	service(route, config, validate)

	logs.Logging.Info(context.Background(), "Service Started!")

	// run the service
	route.Run(fmt.Sprintf(":%d", config.Port))
}

func filenameLog() string {
	var timeNow = time.Now()
	year := timeNow.Year()
	month := timeNow.Month()
	day := timeNow.Day()
	minute := timeNow.Minute()
	second := timeNow.Second()

	return fmt.Sprintf("log_%d%02d%02d%02d%02d.log", year, month, day, minute, second)
}
