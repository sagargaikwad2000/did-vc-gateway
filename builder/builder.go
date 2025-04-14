package builder

import (
	"github.com/sagargaikwad2000/did-vc-gateway/builder/server"
	"github.com/sagargaikwad2000/did-vc-gateway/model/configuration"
)

type Builder struct {
	config configuration.Configuration
	Server server.Server
}

func NewBuilder(config configuration.Configuration) Builder {
	return Builder{config: config}
}

func (b Builder) SetServer() Builder {
	b.Server.SetServer(b.config.Application.Server)
	return b
}

func (b Builder) Build() Builder {
	return b
}
