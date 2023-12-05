package main

import (
	"backend/config"
	"backend/logs"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

const (
	pathConfigEnv = "PATH_CONF"
	fileConfigEnv = "FILE_CONF"
	pathLogEnv    = "PATH_LOG"
)

func main() {

	// get path & file configuration from environment variable
	pathConfig := os.Getenv(pathConfigEnv)
	fileConfig := os.Getenv(fileConfigEnv)
	pathLog := os.Getenv(pathLogEnv)

	// get argument
	argument := os.Args[1]
	createOutputLog, _ := strconv.ParseBool(argument)

	// declare variable
	var err error

	// logging
	var osLog *os.File
	if createOutputLog {

		defer osLog.Close()
		osLog, err = goutil.OpenFile(pathLog, "service.log")
		if err != nil {
			panic(err)
		}

		// override io writer to gin default writer
		gin.DefaultWriter = io.Writer(osLog)
	}

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
	telegram, _ := goutil.NewTele(config.ThirdParty.Telegram.Token, config.ThirdParty.Telegram.ChatId)

	// logs service
	logs.Logging, err = goutil.NewLog(osLog, telegram, createOutputLog)
	if err != nil {
		panic(err)
	}

	defer logs.Logging.Sync()
	defer logs.Logging.Undo()

	// call gin route
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(requestid.New())
	route.Use(goutil.SetContext)
	route.Use(goutil.LogMiddleware(logs.Logging))

	// validation method
	validate, err := goutil.NewValidation()
	if err != nil {
		panic(err)
	}

	// static route
	route.Static("/swagger/", "../docs/openapi")

	// running the service
	service(route, config, validate)

	logs.Logging.Info(context.Background(), fmt.Sprintf("Service Started! at %d", config.Port))

	// run the service
	route.Run(fmt.Sprintf(":%d", config.Port))
}
