package configuration

import (
	"os"
	"strings"

	"github.com/akrck02/valhalla-core/log"
	"github.com/joho/godotenv"
)

type GlobalConfiguration struct {
	Ip     string
	Port   string
	Secret string
	Mongo  string
}

var Params GlobalConfiguration

func LoadConfiguration() {

	err := godotenv.Load(ENV_PATH)

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
	Params = configuration
}

func checkCompulsoryVariables(Configuration GlobalConfiguration) {
	log.Jump()
	log.Line()
	log.Info("Configuration variables")
	log.Line()
	log.Info("IP: " + Configuration.Ip)
	log.Info("PORT: " + Configuration.Port)
	log.Info("SECRET: " + strings.Repeat("*", len(Configuration.Secret)))
	log.Info("MONGO: " + Configuration.Mongo)
	log.Line()
}
