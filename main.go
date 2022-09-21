package main

import (
"github.com/AlanaPeres/Biblioteca/server"
)
func main() {
	s := server.NewServer()
	s.Run()
}