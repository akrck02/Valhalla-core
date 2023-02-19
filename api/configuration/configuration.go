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
}

func LoadConfiguration() GlobalConfiguration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var configuration = GlobalConfiguration{
		Ip:     os.Getenv("IP"),
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
	}

	checkCompulsoryVariables(configuration)
	return configuration
}

func checkCompulsoryVariables(Configuration GlobalConfiguration) {

}
