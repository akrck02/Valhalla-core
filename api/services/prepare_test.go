package services

import (
	"os"
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/log"
)

var setupDone bool = false

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	if setupDone {
		return
	}

	log.Jump()
	log.Info("Setting up test environment...")
	db.SetupTest()
	setupDone = true
	log.Jump()
}
