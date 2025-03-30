package main

import (
	"fmt"

	"github.com/sagargaikwad2000/did-vc-gateway/configuration"
)

func init() {
	config := configuration.MustReadConfig()

	fmt.Printf("config: %+v\n", config)
}

func main() {
	fmt.Println("Hello...")
}
