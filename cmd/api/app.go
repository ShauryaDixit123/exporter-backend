package api

import (
	"context"
	"database/sql"
	"exporterbackend/internal/configs"
	v1 "exporterbackend/internal/handlers/api/v1"
	"exporterbackend/pkg/logging"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/gin-gonic/gin"
)

type app struct {
	engine *gin.Engine
}

func (a *app) run(port int) {

	a.engine.Run(fmt.Sprintf(":%d", port))

	fmt.Printf("Run Http Service on Port :%d", port)
}

func NewPgDbInstance(pgDbConfig configs.PgDbConfig) (*sql.DB, error) {
	var conninfo string = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgDbConfig.Host,
		pgDbConfig.Port,
		pgDbConfig.Username,
		pgDbConfig.Password,
		pgDbConfig.Database,
		pgDbConfig.SSLMode,
	)
	db, err := sql.Open("pgx", conninfo)

	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(pgDbConfig.MaxIdle)
	db.SetMaxOpenConns(pgDbConfig.MaxConnections)

	return db, nil
}

func NewLogger(
	appName configs.AppName,
	logConfig configs.LogConfig,
) (logging.Logger, error) {
	switch logConfig.LogSink {
	case configs.CONSOLE:
		return logging.NewConsoleLogger(
			string(appName),
			logConfig.LogLevel,
		), nil
	case configs.STDOUT:
		return logging.NewStdOutLogger(
			string(appName),
			logConfig.LogLevel,
		), nil
	default:
		return nil, fmt.Errorf(
			"Invalid Log Sink: %v",
			logConfig.LogSink,
		)
	}
}

func NewGoquInstance(pgDB *sql.DB) *goqu.Database {
	dialect := goqu.Dialect("postgres")
	return dialect.DB(pgDB)
}

func NewRequestContext() context.Context {
	return context.WithValue(&gin.Context{}, "start_up", "undefined")
}

func NewHttpEngine(
	v1Routes v1.GroupRoutes,
) *gin.Engine {

	r := gin.New()
	r.RedirectTrailingSlash = true
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	v1Routes.Initialize("/v1", r)

	return r
}

func NewApp(engine *gin.Engine) *app {
	return &app{
		engine: engine,
	}
}

func Initialize(config configs.Config) {
	app, err := InitializeApp(
		config.AppName,
		config.PgDbConfig,
		config.LogConfig,
		config.Context,
	)

	if err != nil {
		panic("There was an error at startup: " + err.Error())
	}

	app.run(config.Port)
}
