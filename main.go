package main

import (
	"github.com/Progsilva/rest_api/server"
)

func main() {
	s := server.NewServer()

	s.Run()
}
