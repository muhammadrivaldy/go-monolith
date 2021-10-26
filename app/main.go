package main

import (
	"backend/models"
	"context"
	"fmt"
	"io"
	"os"

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
	var logs goutil.Logs

	// logging
	var osLog *os.File
	defer osLog.Close()
	osLog, err = goutil.OpenFile("./log", filenameLog())
	if err != nil {
		panic(err)
	}

	// override io writer to gin default writer
	gin.DefaultWriter = io.Writer(osLog)

	// cron for create a log file every day
	c := cron.New()
	c.AddFunc("@midnight", func() {
		osLogNewest, err := goutil.OpenFileNewest(osLog, "./log", filenameLog())
		if err != nil {
			return
		}

		logs.Config(osLogNewest)
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
	var config models.Configuration
	if err := goutil.Configuration(osFile, &config); err != nil {
		panic(err)
	}

	// third-party telegram
	tele, err := goutil.NewTele(config.ThirdParty.Telegram.Token, config.ThirdParty.Telegram.ChatID)
	if err != nil {
		panic(err)
	}

	// logs service
	logs = goutil.NewLog(osLog, tele)
	defer logs.Sync()
	defer logs.Undo()

	logs.Info(context.Background(), "Success Load Configuration")

	// call gin route
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(requestid.New())
	route.Use(goutil.SetContext)
	route.Use(goutil.LogMiddleware(logs))

	// don't remove this code
	route.Static("nominatim", "../nominatim")

	// validation method
	validate, err := goutil.NewValidation()
	if err != nil {
		panic(err)
	}

	// running the service
	service(route, config, logs, validate)

	logs.Info(context.Background(), "Service Started!")

	// run the service
	route.Run(fmt.Sprintf(":%d", config.Port))

}
