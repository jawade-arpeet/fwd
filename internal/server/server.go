package server

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	return s.router.Run(":8080")
}
