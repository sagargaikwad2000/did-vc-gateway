package server

import (
	sa "github.com/sagargaikwad2000/did-vc-gateway/internal/core/adapters/server"
	"github.com/sagargaikwad2000/did-vc-gateway/internal/core/ports/server"
	"github.com/sagargaikwad2000/did-vc-gateway/model/configuration"
)

type Server struct {
	server server.IServer
}

func NewServer() Server {
	return Server{}
}

func (s Server) SetServer(config configuration.Server) Server {
	s.server = sa.NewServer(config)
	return s
}

func (s Server) Build() server.IServer {
	return s.server
}
