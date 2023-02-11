package main

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/services"
	"github.com/akrck02/valhalla-core/utils"
)

func main() {

	logger := utils.Logger
	db.ConnectDatabase(logger)
	services.Start(logger)

}
