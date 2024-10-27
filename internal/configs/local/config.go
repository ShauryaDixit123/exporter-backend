package local

import (
	"context"
	"exporterbackend/internal/configs"
	"exporterbackend/pkg/logging"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetConfig(appName configs.AppName) configs.Config {

	pgPort, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	var apiPort int = 8080
	if len(os.Getenv("PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("PORT"))

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		} else {
			apiPort = p
		}
	}

	var grpcPort int = 5051
	if len(os.Getenv("GRPC_PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("GRPC_PORT"))

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		} else {
			grpcPort = p
		}
	}

	return configs.Config{
		AppName:   appName,
		Port:      apiPort,
		GrpcPort:  grpcPort,
		EnableAPM: false,
		LogConfig: configs.LogConfig{
			LogSink:  configs.CONSOLE,
			LogLevel: logging.INFO,
		},
		PgDbConfig: configs.PgDbConfig{
			Host:           os.Getenv("PG_HOST"),
			Port:           pgPort,
			Username:       os.Getenv("PG_USERNAME"),
			Password:       os.Getenv("PG_PASSWORD"),
			Database:       os.Getenv("PG_DATABASE"),
			DatabasePrefix: os.Getenv("PG_DATABASE_PREFIX"),
			MaxConnections: 5,
			MaxIdle:        1,
			SSLMode:        configs.SSL_MODE_DISABLED,
		},
		Context: context.WithValue(&gin.Context{}, "start_up", "undefined"),
	}
}
