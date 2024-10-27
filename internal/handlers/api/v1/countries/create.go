package countries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi there")
}
