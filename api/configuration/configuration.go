package configuration

import (
	"os"

	"github.com/akrck02/valhalla-core/log"
	"github.com/joho/godotenv"
)

type GlobalConfiguration struct {
	Ip     string
	Port   string
	Secret string
	Mongo  string
}

var Params = LoadConfiguration()

func LoadConfiguration() GlobalConfiguration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var configuration = GlobalConfiguration{
		Ip:     os.Getenv("IP"),
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
		Mongo:  os.Getenv("IP_MONGODB"),
	}

	checkCompulsoryVariables(configuration)
	return configuration
}

func checkCompulsoryVariables(Configuration GlobalConfiguration) {
	log.Info("Checking compulsory variables")
	log.Info("IP: " + Configuration.Ip)
	log.Info("PORT: " + Configuration.Port)
	log.Info("SECRET: " + Configuration.Secret)
	log.Info("MONGO: " + Configuration.Mongo)
}
