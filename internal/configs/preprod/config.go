package preprod

import (
	"exporterbackend/internal/configs"
	"exporterbackend/pkg/logging"
	"fmt"
	"os"
	"strconv"
)

func GetConfig(appName configs.AppName) configs.Config {

	pg_port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return configs.Config{
		AppName:  appName,
		Port:     8080,
		GrpcPort: 5051,
		LogConfig: configs.LogConfig{
			LogSink:  configs.CONSOLE,
			LogLevel: logging.INFO,
		},
		PgDbConfig: configs.PgDbConfig{
			Host:           os.Getenv("PG_HOST"),
			Port:           pg_port,
			Username:       os.Getenv("PG_USERNAME"),
			Password:       os.Getenv("PG_PASSWORD"),
			Database:       os.Getenv("PG_DATABASE"),
			DatabasePrefix: os.Getenv("PG_DATABASE_PREFIX"),
			MaxConnections: 5,
			MaxIdle:        1,
		},
	}
}
