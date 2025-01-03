package v1

import (
	"exporterbackend/internal/handlers/api/v1/countries"
	"exporterbackend/internal/handlers/api/v1/currencies"
	"exporterbackend/internal/handlers/api/v1/orders"
	"exporterbackend/internal/handlers/api/v1/quotes"
	"exporterbackend/internal/handlers/api/v1/users"
	"exporterbackend/internal/handlers/api/v1/workflows"
	"exporterbackend/internal/handlers/api/v1/ws"

	"github.com/gin-gonic/gin"
)

type GroupRoutes interface {
	Initialize(prefix string, r gin.IRouter)
}

type Routes struct {
	countryRoutes     countries.GroupRoutes
	currencyRoutes    currencies.GroupRoutes
	userRoutes        users.GroupRoutes
	workflowRoutes    workflows.GroupRoutes
	orderRoutes       orders.GroupRoutes
	quoteRoutes       quotes.GroupRoutes
	routesMiddlewares RouteMiddlewares
	wsRoutes          ws.GroupRoutes
}

func New(
	countryRoutes countries.GroupRoutes,
	currencyRoutes currencies.GroupRoutes,
	usersRoutes users.GroupRoutes,
	workflowRoutes workflows.GroupRoutes,
	orderRoutes orders.GroupRoutes,
	quoteRoutes quotes.GroupRoutes,
	routesMiddlewares RouteMiddlewares,
	wsRoutes ws.GroupRoutes,
) *Routes {
	return &Routes{
		countryRoutes:     countryRoutes,
		currencyRoutes:    currencyRoutes,
		userRoutes:        usersRoutes,
		workflowRoutes:    workflowRoutes,
		orderRoutes:       orderRoutes,
		quoteRoutes:       quoteRoutes,
		routesMiddlewares: routesMiddlewares,
		wsRoutes:          wsRoutes,
	}
}

// func withUserDetails(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		accountId := "12345" // Example data, typically extracted from headers/tokens
// 		userId := "67890"

// 		ctx := context.WithValue(r.Context(), "accountIdKey", accountId)
// 		ctx = context.WithValue(ctx, "userIdKey", userId)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func (ro *Routes) Initialize(prefix string, r gin.IRouter) {
	v1 := r.Group(prefix)
	{
		ro.userRoutes.Initialize("/users", v1)
		ro.countryRoutes.Initialize("/countries", v1)
		ro.currencyRoutes.Initialize("/currencies", v1)
		ro.wsRoutes.Initialize("/ws", v1)
	}
	v1.Use(ro.routesMiddlewares.PermissionsMiddleware())
	{

		ro.workflowRoutes.Initialize("/workflows", v1)
		ro.orderRoutes.Initialize("/orders", v1)
		ro.quoteRoutes.Initialize("/quotes", v1)
	}
}
