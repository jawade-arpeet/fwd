package server

import (
	"fmt"
	"fwd/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(cfg *config.ServerConfig) *Server {
	return &Server{
		config: cfg,
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	return s.router.Run(addr)
}
