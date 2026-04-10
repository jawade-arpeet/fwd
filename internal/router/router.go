package router

import (
	"fwd/internal/constants"
	"fwd/internal/handler"
	v1 "fwd/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter(env constants.Env, handler *handler.Handler) *gin.Engine {
	if env == constants.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	v1.MountV1Routes(router, handler)

	return router
}
