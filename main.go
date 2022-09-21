package main

import (
	"github.com/AlanaPeres/Biblioteca/database"
	"github.com/AlanaPeres/Biblioteca/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()
}