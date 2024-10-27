package currencies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi there")
}
