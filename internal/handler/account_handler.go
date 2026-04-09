package handler

import (
	"fwd/internal/errs"
	"fwd/internal/model"
	"fwd/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) SignUp(ctx *gin.Context) {
	var payload *model.SignUpPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		zap.L().Error(
			"error binding payload",
			zap.String("operation", "AccountHandler.SignUp"),
			zap.Error(err),
		)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": errs.ErrInvalidPayload.Error()},
		)
		return
	}

	if err := h.accountService.SignUp(ctx.Request.Context(), payload); err != nil {
		switch err {
		case errs.ErrAccountAlreadyExists:
			ctx.JSON(
				http.StatusConflict,
				gin.H{"error": err.Error()},
			)
		default:
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": errs.ErrInternalServer.Error()},
			)
		}
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{"message": "sign-up successful"},
	)
}

func (h *AccountHandler) SignIn(ctx *gin.Context) {
	var payload *model.SignInPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		zap.L().Error(
			"error binding payload",
			zap.String("operation", "AccountHandler.SignIn"),
			zap.Error(err),
		)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": errs.ErrInvalidPayload.Error()},
		)
		return
	}

	if err := h.accountService.SignIn(
		ctx.Request.Context(), payload,
	); err != nil {
		switch err {
		case errs.ErrAccountDoesNotExists:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errs.ErrInvalidPassword:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": errs.ErrInternalServer.Error()},
			)
		}
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "sign-in successful"},
	)
}
