package server

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
	log.Print("o servidor está passando na porta: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}