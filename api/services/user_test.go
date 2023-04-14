package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
)

func TestRegisterMinimumCharacters(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    "testingapi@testing.org",
		Password: "password",
		Username: "test-user",
	}

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with less than minimum characters")
		return
	}

	if err.Error != error.SHORT_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}
}

func TestRegisterSpecialCharacters(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    "testingapi@testing.org",
		Password: "Passwordpassword1",
		Username: "test-user",
	}

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any special character")
		return
	}

	if err.Error != error.NO_SPECIAL_CHARACTERS_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}
}

func TestRegisterUpperCaseLoweCase(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    "testingapi@testing.org",
		Password: "passwordpassword1#",
		Username: "test-user",
	}
	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	user = &models.User{
		Email:    "testingapi@testing.org",
		Password: "PASSWORDPASSWORD1#",
		Username: "test-user",
	}
	err = Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}
}
