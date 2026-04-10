package handler

import (
	"fwd/internal/errs"
	"fwd/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlatformHandler struct {
	platformService *service.PlatformService
}

func NewPlatformHandler(
	platformService *service.PlatformService,
) *PlatformHandler {
	return &PlatformHandler{
		platformService: platformService,
	}
}

func (h *PlatformHandler) GetAllPlatforms(ctx *gin.Context) {
	platforms, err := h.platformService.GetAllPlatforms(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, platforms)
}
