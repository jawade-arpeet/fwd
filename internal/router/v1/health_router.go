package v1

import (
	"fwd/internal/handler"

	"github.com/gin-gonic/gin"
)

func MountHealthRoutes(router *gin.RouterGroup, handler *handler.HealthHandler) {
	health := router.Group("/health")
	{
		health.GET("/", handler.Check)
	}
}
