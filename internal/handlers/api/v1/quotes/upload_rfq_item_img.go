package quotes

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadRFQItemImg(ctx *gin.Context) {
	var imgs []rdbms.CreateImage
	if er := ctx.ShouldBindJSON(&imgs); er != nil {
		ctx.JSON(500, er)
	}
	if res, er := h.imagesService.GetSignedURLAndSave(imgs); er != nil {
		ctx.JSON(500, er)
	} else {
		ctx.JSON(200, res)
	}
}
