package server

import (
	"fmt"
	"fwd/internal/config"
	"fwd/internal/handler"
	"fwd/internal/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(cfg *config.ServerConfig) *Server {
	handler := handler.NewHandler()
	router := router.NewRouter(handler)

	return &Server{
		config: cfg,
		router: router,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	return s.router.Run(addr)
}
