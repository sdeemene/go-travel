package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database    Database
	Server      Server
	Environment Environment
	Jwt         Jwt
}

type Database struct {
	LocalUri     string
	DatabaseName string
	RemoteUri    string
}

type Server struct {
	Port string
}
type Environment struct {
	Profile string
}

type Jwt struct {
	AccessKey  string
	RefreshKey string
	Issuer     string
}

func GetConfiguration() Configuration {

	conf := Configuration{}

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file, %s\n", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Error decoding config, %v\n", err)
	}
	return conf
}
