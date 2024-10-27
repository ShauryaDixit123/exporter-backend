package countries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Read(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi there")
}
