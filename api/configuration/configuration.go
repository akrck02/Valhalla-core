package configuration

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/akrck02/valhalla-core/log"
	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	BASE_PATH  = filepath.Dir(b) + "/../"
)

type GlobalConfiguration struct {
	Ip     string
	Port   string
	Secret string
	Mongo  string
}

var Params = LoadConfiguration()

func LoadConfiguration() GlobalConfiguration {
	err := godotenv.Load(BASE_PATH + ".env")

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
