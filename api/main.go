package main

import (
	"github.com/akrck02/valhalla-core/services"
	"github.com/akrck02/valhalla-core/utils"
)

func main() {

	logger := utils.Logger
	services.Start(logger)
}
