package main

import (
	"github.com/sagargaikwad2000/did-vc-gateway/builder"
	"github.com/sagargaikwad2000/did-vc-gateway/configuration"
)

// func init() {
// 	config := configuration.MustReadConfig()

// 	fmt.Printf("config: %+v\n", config)
// }

func main() {

	config := configuration.MustReadConfig()

	b := builder.NewBuilder(*config).SetServer().Build()

	s := b.Server.SetServer(config.Application.Server).Build()
	s.MustStart()
}
