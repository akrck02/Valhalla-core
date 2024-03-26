package services

import (
	"log"
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/mock"
	"github.com/akrck02/valhalla-core/models"
)

func TestDeviceExists(t *testing.T) {

	// Connect database
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create user
	user := models.User{
		Email:    mock.Email(),
		Username: mock.Username(),
		Password: mock.Password(),
	}

	err := Register(conn, client, &user)

	if err != nil {
		t.Error(err)
	}

	// add device to user
	var expected = models.Device{
		Token:     mock.Token(),
		User:      user.Email,
		Address:   mock.Ip(),
		UserAgent: mock.Platform(),
	}

	_, error := AddUserDevice(conn, client, &user, &expected)

	if error != nil {
		t.Error(err)
	}

	// check if device exists
	coll := client.Database(db.CurrentDatabase).Collection(db.DEVICE)
	obtained, error := FindDevice(conn, coll, &expected)

	if error != nil {
		t.Error(err)
	}

	if obtained == nil {
		t.Error("Device not found")
	}

	log.Print("Device expected: ", expected)
	log.Print("Device found: ", obtained)

	// delete device
	error = DeleteDevice(conn, client, &expected)

	if error != nil {
		t.Error(err)
	}

	// delete user
	err = DeleteUser(conn, client, &user)

	if err != nil {
		t.Error(err)
	}

}

func TestDeviceNotExists(t *testing.T) {

	// Connect database
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create user
	user := models.User{
		Email:    mock.Email(),
		Username: mock.Username(),
		Password: mock.Password(),
	}

	err := Register(conn, client, &user)

	if err != nil {
		t.Error(err)
	}

	// add device to user
	var expected = models.Device{
		Token:     mock.Token(),
		User:      user.Email,
		Address:   mock.Ip(),
		UserAgent: mock.Platform(),
	}

	_, error := AddUserDevice(conn, client, &user, &expected)

	if error != nil {
		t.Error(err)
	}

	// check if device exists
	coll := client.Database(db.CurrentDatabase).Collection(db.DEVICE)
	obtained, error := FindDevice(conn, coll, &models.Device{
		Token: mock.Token(),
	})

	if error == nil || obtained != nil {
		t.Error("Device found")
	}

	log.Print("Device expected: ", expected)
	log.Print("Device not found: ", obtained)

	// delete device
	error = DeleteDevice(conn, client, &expected)

	if error != nil {
		t.Error(err)
	}

	// delete user
	err = DeleteUser(conn, client, &user)

	if err != nil {
		t.Error(err)
	}

}
