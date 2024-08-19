package main

import (
	"backend/config"
	"backend/logs"
	"backend/tracer"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
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

	// get parameters
	outputLogArgument := os.Args[1]
	filenameLogArgument := os.Args[2]
	serviceDo := os.Args[3]

	createOutputLog, _ := strconv.ParseBool(outputLogArgument)

	// declare variable
	var err error

	// logging
	var osLog *os.File
	if createOutputLog {

		defer osLog.Close()
		osLog, err = goutil.OpenFile(pathLog, filenameLogArgument)
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
	var configuration config.Configuration
	if err := goutil.Configuration(osFile, &configuration); err != nil {
		panic(err)
	}

	// third-party telegram
	telegram, _ := goutil.NewTele(configuration.ThirdParty.Telegram.Token, configuration.ThirdParty.Telegram.ChatID)

	// logs service
	logs.Logging, err = goutil.NewLog(osLog, telegram, createOutputLog)
	if err != nil {
		panic(err)
	}

	defer logs.Logging.Sync()
	defer logs.Logging.Undo()

	// set config to public variable
	config.Config = configuration

	switch serviceDo {
	case "run-service":
		runService(configuration)
	case "run-migration":
		runMigration(configuration)
	default:
		panic("unexpected service do")
	}
}

func runMigration(config config.Configuration) {

	databaseMigration(config)
}

func runService(config config.Configuration) {

	// declare variable
	var err error

	// run apm
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(config.Uptrace.DSN),
		uptrace.WithServiceName(config.Uptrace.ServiceName),
		uptrace.WithDeploymentEnvironment(config.Uptrace.Environment),
	)

	defer uptrace.Shutdown(context.Background())

	// set up tracer
	tracer.Tracer = otel.Tracer(fmt.Sprintf("%s %s", config.Uptrace.ServiceName, config.Uptrace.Environment))

	// cors set up
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}

	// call gin route
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(cors.New(corsConfig))
	route.Use(requestid.New())
	route.Use(otelgin.Middleware(config.Uptrace.ServiceName))
	route.Use(goutil.SetContext)
	route.Use(goutil.LogMiddleware(logs.Logging))

	// validation method
	validate, err := goutil.NewValidation()
	if err != nil {
		panic(err)
	}

	// static route
	route.StaticFile("/privacy", "./../privacy_policy.html")
	route.Static(fmt.Sprintf("/swagger/%s", config.SwaggerAPIKey), "../docs/openapi")
	route.GET("/.well-known/assetlinks.json", func(ctx *gin.Context) {
		ctx.File("../assets/assetlinks.json")
	})

	// running the service
	service(route, config, validate)

	logs.Logging.Info(context.Background(), fmt.Sprintf("Service Started! at %d", config.Port))

	// run the service
	route.Run(fmt.Sprintf(":%d", config.Port))
}
