//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package api

import (
	"context"
	"exporterbackend/internal/configs"
	"exporterbackend/internal/core/ports"
	"exporterbackend/internal/core/services/countriessrv"
	"exporterbackend/internal/core/services/currenciessrv"
	"exporterbackend/internal/core/services/orderssrv"
	"exporterbackend/internal/core/services/userssrv"
	"exporterbackend/internal/core/services/workflowssrv"
	v1 "exporterbackend/internal/handlers/api/v1"
	"exporterbackend/internal/handlers/api/v1/countries"
	"exporterbackend/internal/handlers/api/v1/currencies"
	"exporterbackend/internal/handlers/api/v1/orders"
	"exporterbackend/internal/handlers/api/v1/users"
	"exporterbackend/internal/handlers/api/v1/workflows"
	"exporterbackend/internal/repositories/pgdb/accountsrepo"
	"exporterbackend/internal/repositories/pgdb/countriesrepo"
	"exporterbackend/internal/repositories/pgdb/currenciesrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/lineitemsrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/purchaseorderrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/salesorderrepo"
	"exporterbackend/internal/repositories/pgdb/usersrepo"
	"exporterbackend/internal/repositories/pgdb/workflowrepo"

	"github.com/google/wire"
)

func InitializeApp(
	appName configs.AppName,
	pgDbConfig configs.PgDbConfig,
	logConfig configs.LogConfig,
	context context.Context,
) (*app, error) {
	wire.Build(
		NewLogger,
		NewPgDbInstance,
		NewGoquInstance,

		//Repositories
		countriesrepo.New,
		currenciesrepo.New,
		usersrepo.New,
		accountsrepo.New,
		workflowrepo.New,
		purchaseorderrepo.New,
		lineitemsrepo.New,
		salesorderrepo.New,
		//Repo Bindings
		wire.Bind(new(ports.RdbmsCountriesRepository), new(*countriesrepo.Repository)),
		wire.Bind(new(ports.RdbmsCurrenciesRepository), new(*currenciesrepo.Repository)),
		wire.Bind(new(ports.RdbmsUsersRepository), new(*usersrepo.Repository)),
		wire.Bind(new(ports.RdbmsAccountsRepository), new(*accountsrepo.Repository)),
		wire.Bind(new(ports.RdbmsWorkflowRepository), new(*workflowrepo.Repository)),
		wire.Bind(new(ports.RdbmsPurchaseOrderRepoistory), new(*purchaseorderrepo.Repository)),
		wire.Bind(new(ports.RdbmsSalesOrderRepoistory), new(*salesorderrepo.Repository)),
		wire.Bind(new(ports.RdbmsPurchaseOrderLineItemsRepoistory), new(*lineitemsrepo.Repository)),
		//Services
		countriessrv.New,
		currenciessrv.New,
		userssrv.New,
		workflowssrv.New,
		orderssrv.New,
		//Service Bindings
		wire.Bind(new(ports.CountriesService), new(*countriessrv.Service)),
		wire.Bind(new(ports.CurrenciesService), new(*currenciessrv.Service)),
		wire.Bind(new(ports.UsersService), new(*userssrv.Service)),
		wire.Bind(new(ports.WorkflowService), new(*workflowssrv.Service)),
		wire.Bind(new(ports.OrdersService), new(*orderssrv.Service)),
		//RouteHandlers
		countries.NewHandler,
		currencies.NewHandler,
		users.NewHandler,
		workflows.NewHandler,
		orders.NewHandler,
		//RouteHandlerBindings
		wire.Bind(new(countries.RoutesHandler), new(*countries.Handler)),
		wire.Bind(new(currencies.RoutesHandler), new(*currencies.Handler)),
		wire.Bind(new(users.RoutesHandler), new(*users.Handler)),
		wire.Bind(new(workflows.RoutesHandler), new(*workflows.Handler)),
		wire.Bind(new(orders.RoutesHandler), new(*orders.Handler)),

		//Group Routes
		countries.New,
		currencies.New,
		users.New,
		workflows.New,
		orders.New,
		v1.New,

		//Group Route Bindings
		wire.Bind(new(countries.GroupRoutes), new(*countries.Routes)),
		wire.Bind(new(currencies.GroupRoutes), new(*currencies.Routes)),
		wire.Bind(new(users.GroupRoutes), new(*users.Routes)),
		wire.Bind(new(workflows.GroupRoutes), new(*workflows.Routes)),
		wire.Bind(new(orders.GroupRoutes), new(*orders.Routes)),

		wire.Bind(new(v1.GroupRoutes), new(*v1.Routes)),

		NewHttpEngine,
		NewApp,
	)
	return &app{}, nil
}