package server

import (
	"fmt"
	"fwd/internal/client"
	"fwd/internal/config"
	"fwd/internal/handler"
	"fwd/internal/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(cfg *config.ServerConfig, client *client.Client) *Server {
	handler := handler.NewHandler()
	router := router.NewRouter(handler)

	return &Server{
		config: cfg,
		router: router,
	}
}

func (s *Server) Start() error {
	zap.L().Info("starting server", zap.String("port", s.config.Port))
	addr := fmt.Sprintf(":%s", s.config.Port)
	return s.router.Run(addr)
}
