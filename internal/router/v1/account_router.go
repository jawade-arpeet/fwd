package v1

import (
	"fwd/internal/handler"

	"github.com/gin-gonic/gin"
)

func MountAccountRoutes(
	router *gin.RouterGroup,
	handler *handler.AccountHandler,
) {
	account := router.Group("/account")
	{
		account.POST("/sign-up", handler.SignUp)
		account.POST("/sign-in", handler.SignIn)
	}
}
