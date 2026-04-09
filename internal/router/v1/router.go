package v1

import (
	"fwd/internal/handler"

	"github.com/gin-gonic/gin"
)

func MountV1Routes(router *gin.Engine, handler *handler.Handler) {
	v1 := router.Group("/api/v1")

	MountHealthRoutes(v1, handler.Health)
	MountAccountRoutes(v1, handler.Account)
}
