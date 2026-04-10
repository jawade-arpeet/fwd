package v1

import (
	"fwd/internal/handler"

	"github.com/gin-gonic/gin"
)

func MountPlatformRoutes(router *gin.RouterGroup, handler *handler.PlatformHandler) {
	platform := router.Group("/platform")
	{
		platform.GET("", handler.GetAllPlatforms)
	}
}
