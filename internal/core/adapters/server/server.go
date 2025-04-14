package server

import (
	"fmt"

	"github.com/sagargaikwad2000/did-vc-gateway/model/configuration"
)

type Server struct {
	Port string
	Host string
}

func NewServer(config configuration.Server) Server {
	return Server{
		Port: config.Port,
		Host: config.Host,
	}
}

/*
Start the server
*/
func (s Server) MustStart() {
	fmt.Println("Server starting...", s.Host, s.Port)
}

/*
Stop the server
*/
func (s Server) MustStop() {}
