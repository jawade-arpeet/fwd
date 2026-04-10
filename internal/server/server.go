package server

import (
	"fmt"
	"fwd/internal/client"
	"fwd/internal/config"
	"fwd/internal/handler"
	"fwd/internal/repo"
	"fwd/internal/router"
	"fwd/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(cfg *config.ServerConfig, client *client.Client) *Server {
	repo := repo.NewRepo(client)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := router.NewRouter(cfg.RunEnv, handler)

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
