package server

import (
	"log"

	"example.com/go-auth-globo/internal/route"
	"example.com/go-auth-globo/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	c := NewContainer()

	if err := c.Invoke(route.ConfigRoute); err != nil {
		panic(err)
	}
	log.Println("Server running on port: ", s.port)
	if err := c.Invoke(startHttpServer); err != nil {
		panic(err)
	}
}

func startHttpServer(e *gin.Engine) {
	service.Logger().Info("BINDING ADDRESS...")

	e.Run(":5000")
}
