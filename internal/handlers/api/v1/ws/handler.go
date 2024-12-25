package ws

import (
	"exporterbackend/pkg/logging"
	"exporterbackend/pkg/socket"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger     logging.Logger
	socketPool *map[string]*socket.Client
}

func NewHandler(logger logging.Logger, socketPool *map[string]*socket.Client) *Handler {
	return &Handler{
		logger:     logger,
		socketPool: socketPool,
	}
}

type RoutesHandler interface {
	Init(ctx *gin.Context)
}

func (h *Handler) Init(ctx *gin.Context) {
	h.serveWs(ctx.Writer, ctx.Request)
}

func (h *Handler) serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("socket Endpoint Hit", r.URL.Query().Get("socket_type"), r.URL.Query().Get("d"))
	localPool := socket.NewPool()
	id := r.URL.Query().Get("id")
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &socket.Client{
		ID:   id,
		Conn: conn,
		Pool: localPool,
	}
	switch r.URL.Query().Get("socket_type") {
	case "notification":
		(*h.socketPool)[fmt.Sprintf("notification_%s", id)] = client
	case "chat_pool":
		(*h.socketPool)[r.Header.Get("chat_pool_id")] = client
	case "":
		(*h.socketPool)[id] = client
	default:
		(*h.socketPool)[id] = client
	}
	// serverPool.Register <- client
	// client.Read()
}
