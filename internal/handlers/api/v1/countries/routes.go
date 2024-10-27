package countries

import (
	"github.com/gin-gonic/gin"
)

type GroupRoutes interface {
	Initialize(prefix string, r gin.IRouter)
}

type Routes struct {
	handler RoutesHandler
}

func New(
	handler RoutesHandler,
) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (ro *Routes) Initialize(prefix string, r gin.IRouter) {

	g := r.Group(prefix)
	{
		g.GET("", ro.handler.List)
		g.GET("/:id", ro.handler.Read)
		g.POST("", ro.handler.Create)
		g.PATCH("/:id", ro.handler.Update)
		g.PATCH("/:id/activate", ro.handler.Activate)
		g.PATCH("/:id/deactivate", ro.handler.Deactivate)
		g.DELETE("/:id", ro.handler.Delete)
	}
}
