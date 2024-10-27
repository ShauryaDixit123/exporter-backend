//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package rpc

import (
	"exporterbackend/internal/configs"
	"exporterbackend/internal/core/ports"
	"exporterbackend/internal/core/services/countriessrv"
	"exporterbackend/internal/core/services/currenciessrv"
	"exporterbackend/internal/handlers/rpc/v1/countries"
	"exporterbackend/internal/handlers/rpc/v1/currencies"
	"exporterbackend/internal/repositories/pgdb/countriesrepo"
	"exporterbackend/internal/repositories/pgdb/currenciesrepo"

	pb "exporterbackend/proto/gen/app/v1"

	"github.com/google/wire"
)

func InitializeApp(
	appName configs.AppName,
	pgDbConfig configs.PgDbConfig,
	logConfig configs.LogConfig,
) (*app, error) {
	wire.Build(
		NewLogger,
		NewPgDbInstance,
		NewGoquInstance,

		//Repositories
		countriesrepo.New,
		currenciesrepo.New,

		//Repo Bindings
		wire.Bind(new(ports.RdbmsCountriesRepository), new(*countriesrepo.Repository)),
		wire.Bind(new(ports.RdbmsCurrenciesRepository), new(*currenciesrepo.Repository)),

		//Services
		countriessrv.New,
		currenciessrv.New,

		//Service Bindings
		wire.Bind(new(ports.CountriesService), new(*countriessrv.Service)),
		wire.Bind(new(ports.CurrenciesService), new(*currenciessrv.Service)),

		//RouteHandlers
		countries.NewHandler,
		currencies.NewHandler,

		//RouteHandlerBindings
		wire.Bind(new(pb.CountriesServiceServer), new(*countries.Handler)),
		wire.Bind(new(pb.CurrenciesServiceServer), new(*currencies.Handler)),

		NewGrpcServer,
		NewApp,
	)
	return &app{}, nil
}
