package configuration

import (
	"log"
	"os"

	"github.com/sagargaikwad2000/did-vc-gateway/model/configuration"
	"github.com/spf13/viper"
)

func MustReadConfig() *configuration.Configuration {

	// Read the config file
	viper.SetConfigFile(os.Getenv("CONFIG"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading config:", err)
	}

	// Set value
	var config configuration.Configuration

	config.Application.Server.Host = viper.GetString("application.server.host")
	config.Application.Server.Port = viper.GetInt("application.server.port")

	return &config
}
