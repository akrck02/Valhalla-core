package main

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/services"
	"github.com/akrck02/valhalla-core/utils"
)

func main() {

	logger := utils.Logger
	client := db.CreateClient(logger)
	conn := db.Connect(logger, *client)
	services.Start(logger, conn, *client)
	db.Disconnect(logger, *client, conn)

}
