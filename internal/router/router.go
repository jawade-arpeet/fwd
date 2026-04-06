package router

import (
	"fwd/internal/handler"
	v1 "fwd/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler) *gin.Engine {
	router := gin.Default()

	v1.MountV1Routes(router, handler)

	return router
}
