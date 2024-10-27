package rpc

import (
	"database/sql"
	"exporterbackend/internal/configs"
	"exporterbackend/pkg/logging"
	pb "exporterbackend/proto/gen/app/v1"
	"fmt"
	"log"
	"net"

	"github.com/doug-martin/goqu/v9"
	"google.golang.org/grpc"
)

type app struct {
	grpcServer *grpc.Server
}

func (a *app) run(port int) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := a.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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

func NewGrpcServer(
	currenciesHandler pb.CurrenciesServiceServer,
	countriesHandler pb.CountriesServiceServer,
) *grpc.Server {
	s := grpc.NewServer()

	pb.RegisterCountriesServiceServer(s, countriesHandler)
	pb.RegisterCurrenciesServiceServer(s, currenciesHandler)

	return s
}

func NewApp(grpcServer *grpc.Server) *app {
	return &app{
		grpcServer: grpcServer,
	}
}

func Initialize(config configs.Config) {
	app, err := InitializeApp(
		config.AppName,
		config.PgDbConfig,
		config.LogConfig,
	)

	if err != nil {
		panic("There was an error at startup: " + err.Error())
	}

	app.run(config.GrpcPort)
}
