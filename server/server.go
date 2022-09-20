package server

import (
	"github.com/gin-gonic/gin"
	"github.com/AlanaPeres/Biblioteca/tree/main/server/routes"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}