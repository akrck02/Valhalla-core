package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
)

func TestCreateTeam(t *testing.T) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

}
