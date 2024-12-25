//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package api

import (
	"context"
	"exporterbackend/internal/common/helper"
	"exporterbackend/internal/common/helper/aws"
	"exporterbackend/internal/configs"
	"exporterbackend/internal/core/ports"
	"exporterbackend/internal/core/services/countriessrv"
	"exporterbackend/internal/core/services/currenciessrv"
	"exporterbackend/internal/core/services/imagessrv"
	"exporterbackend/internal/core/services/orderssrv"
	"exporterbackend/internal/core/services/quotessrv"
	"exporterbackend/internal/core/services/userssrv"
	"exporterbackend/internal/core/services/workflowssrv"
	v1 "exporterbackend/internal/handlers/api/v1"
	"exporterbackend/internal/handlers/api/v1/countries"
	"exporterbackend/internal/handlers/api/v1/currencies"
	"exporterbackend/internal/handlers/api/v1/orders"
	"exporterbackend/internal/handlers/api/v1/quotes"
	"exporterbackend/internal/handlers/api/v1/users"
	"exporterbackend/internal/handlers/api/v1/workflows"
	"exporterbackend/internal/handlers/api/v1/ws"
	"exporterbackend/internal/repositories/pgdb/accountsrepo"
	"exporterbackend/internal/repositories/pgdb/countriesrepo"
	"exporterbackend/internal/repositories/pgdb/currenciesrepo"
	"exporterbackend/internal/repositories/pgdb/imagesrepo"
	"exporterbackend/internal/repositories/pgdb/locationsrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/lineitemsrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/purchaseorderrepo"
	"exporterbackend/internal/repositories/pgdb/ordersrepo/salesorderrepo"
	"exporterbackend/internal/repositories/pgdb/quotesrepo"
	"exporterbackend/internal/repositories/pgdb/rolesrepo"
	"exporterbackend/internal/repositories/pgdb/usersrepo"
	"exporterbackend/internal/repositories/pgdb/workflowrepo"

	"github.com/google/wire"
)

func InitializeApp(
	appName configs.AppName,
	pgDbConfig configs.PgDbConfig,
	logConfig configs.LogConfig,
	context context.Context,
	s3Cofig configs.S3Config,
) (*app, error) {
	wire.Build(
		NewLogger,
		NewPgDbInstance,
		NewGoquInstance,
		NewS3Session,
		NewSocketPoolMap,
		//Repositories
		helper.NewHelperRepository,
		v1.NewMiddleware,
		countriesrepo.New,
		currenciesrepo.New,
		usersrepo.New,
		accountsrepo.New,
		workflowrepo.New,
		purchaseorderrepo.New,
		lineitemsrepo.New,
		salesorderrepo.New,
		quotesrepo.New,
		rolesrepo.New,
		locationsrepo.New,
		aws.NewS3,
		imagesrepo.New,
		//Repo Bindings
		wire.Bind(new(aws.S3Service), new(*aws.S3)),
		wire.Bind(new(helper.HelperFunctions), new(*helper.HelperRepository)),
		wire.Bind(new(v1.RouteMiddlewares), new(*v1.RouteMiddleware)),
		wire.Bind(new(ports.RdbmsCountriesRepository), new(*countriesrepo.Repository)),
		wire.Bind(new(ports.RdbmsCurrenciesRepository), new(*currenciesrepo.Repository)),
		wire.Bind(new(ports.RdbmsUsersRepository), new(*usersrepo.Repository)),
		wire.Bind(new(ports.RdbmsAccountsRepository), new(*accountsrepo.Repository)),
		wire.Bind(new(ports.RdbmsWorkflowRepository), new(*workflowrepo.Repository)),
		wire.Bind(new(ports.RdbmsPurchaseOrderRepoistory), new(*purchaseorderrepo.Repository)),
		wire.Bind(new(ports.RdbmsSalesOrderRepoistory), new(*salesorderrepo.Repository)),
		wire.Bind(new(ports.RdbmsPurchaseOrderLineItemsRepoistory), new(*lineitemsrepo.Repository)),
		wire.Bind(new(ports.RdbmsQuotesRepository), new(*quotesrepo.Repository)),
		wire.Bind(new(ports.RdbmsRolesRepository), new(*rolesrepo.Repository)),
		wire.Bind(new(ports.RdbmsLocationsRepository), new(*locationsrepo.Repository)),
		wire.Bind(new(ports.RdbmsImagesRepository), new(*imagesrepo.Repository)),
		//Services
		countriessrv.New,
		currenciessrv.New,
		userssrv.New,
		workflowssrv.New,
		orderssrv.New,
		quotessrv.New,
		imagessrv.New,
		//Service Bindings
		wire.Bind(new(ports.CountriesService), new(*countriessrv.Service)),
		wire.Bind(new(ports.CurrenciesService), new(*currenciessrv.Service)),
		wire.Bind(new(ports.UsersService), new(*userssrv.Service)),
		wire.Bind(new(ports.WorkflowService), new(*workflowssrv.Service)),
		wire.Bind(new(ports.OrdersService), new(*orderssrv.Service)),
		wire.Bind(new(ports.QuotesService), new(*quotessrv.Service)),
		wire.Bind(new(ports.ImagesService), new(*imagessrv.Service)),

		//RouteHandlers
		countries.NewHandler,
		currencies.NewHandler,
		users.NewHandler,
		workflows.NewHandler,
		orders.NewHandler,
		quotes.NewHandler,
		ws.NewHandler,
		//RouteHandlerBindings
		wire.Bind(new(countries.RoutesHandler), new(*countries.Handler)),
		wire.Bind(new(currencies.RoutesHandler), new(*currencies.Handler)),
		wire.Bind(new(users.RoutesHandler), new(*users.Handler)),
		wire.Bind(new(workflows.RoutesHandler), new(*workflows.Handler)),
		wire.Bind(new(orders.RoutesHandler), new(*orders.Handler)),
		wire.Bind(new(quotes.RoutesHandler), new(*quotes.Handler)),
		wire.Bind(new(ws.RoutesHandler), new(*ws.Handler)),
		//Group Routes
		countries.New,
		currencies.New,
		users.New,
		workflows.New,
		orders.New,
		quotes.New,
		ws.New,
		v1.New,

		//Group Route Bindings
		wire.Bind(new(countries.GroupRoutes), new(*countries.Routes)),
		wire.Bind(new(currencies.GroupRoutes), new(*currencies.Routes)),
		wire.Bind(new(users.GroupRoutes), new(*users.Routes)),
		wire.Bind(new(workflows.GroupRoutes), new(*workflows.Routes)),
		wire.Bind(new(orders.GroupRoutes), new(*orders.Routes)),
		wire.Bind(new(quotes.GroupRoutes), new(*quotes.Routes)),
		wire.Bind(new(ws.GroupRoutes), new(*ws.Routes)),
		wire.Bind(new(v1.GroupRoutes), new(*v1.Routes)),

		NewHttpEngine,
		NewApp,
	)
	return &app{}, nil
}
