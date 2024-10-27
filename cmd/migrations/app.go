package migrations

import (
	"database/sql"
	"embed"
	"exporterbackend/internal/configs"
	"exporterbackend/pkg/logging"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

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

func Initialize(embedMigrations embed.FS, config configs.Config) {

	logger, err := NewLogger(config.AppName, config.LogConfig)

	if err != nil {
		panic(err)
	}

	pgdbConn, err := NewPgDbInstance(config.PgDbConfig)

	if err != nil {
		logger.Critical(
			"POSTGRES_STARTUP_FAILED",
			"Unable to connect to postgresql",
			err,
			map[string]any{},
			map[string]any{
				"host":     config.PgDbConfig.Host,
				"port":     config.PgDbConfig.Port,
				"username": config.PgDbConfig.Username,
				"database": config.PgDbConfig.Database,
				"sslMode":  config.PgDbConfig.SSLMode,
			},
		)
		panic(err)
	}

	// setup database
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Critical(
			"GOOSE_STARTUP_ISSUES",
			"Unable to set Goose dialect to postgres",
			err,
			map[string]any{},
			map[string]any{
				"host":     config.PgDbConfig.Host,
				"port":     config.PgDbConfig.Port,
				"username": config.PgDbConfig.Username,
				"database": config.PgDbConfig.Database,
				"sslMode":  config.PgDbConfig.SSLMode,
			},
		)
		panic(err)
	}

	if err := goose.Up(pgdbConn, "migrations"); err != nil {
		logger.Critical(
			"GOOSE_MIGRATION_FAILED",
			"Migration up failed with an error",
			err,
			map[string]any{},
			map[string]any{
				"host":     config.PgDbConfig.Host,
				"port":     config.PgDbConfig.Port,
				"username": config.PgDbConfig.Username,
				"database": config.PgDbConfig.Database,
			},
		)
		panic(err)
	}
}
