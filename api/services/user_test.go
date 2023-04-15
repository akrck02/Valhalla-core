package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
)

func TestRegister(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    "testingapiregister@testing.org",
		Password: "PasswordPassword1#",
		Username: "test-user",
	}

	err := Register(conn, client, &user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// delete the user
	err = DeleteUser(conn, client, &user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

}

func TestRegisterMinimumCharacters(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    "testingapi@testing.org",
		Password: "password",
		Username: "test-user",
	}

	err := Register(conn, client, &user)

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

	var user = models.User{
		Email:    "testingapi@testing.org",
		Password: "Passwordpassword1",
		Username: "test-user",
	}

	err := Register(conn, client, &user)

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

	var user = models.User{
		Email:    "testingapi@testing.org",
		Password: "passwordpassword1#",
		Username: "test-user",
	}
	err := Register(conn, client, &user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	user = models.User{
		Email:    "testingapi@testing.org",
		Password: "PASSWORDPASSWORD1#",
		Username: "test-user",
	}
	err = Register(conn, client, &user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}
}

func TestChangeUserPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    "testingapipasswordchange@testing.org",
		Password: "PasswordPassword1#",
		Username: "test-user",
	}

	err := Register(conn, client, &user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	user.Password = "Passwordpassword1#"

	// change the user password
	err = ChangePassword(conn, client, &user)

	if err != nil {
		t.Error("The user password was not changed")
		return
	}

	// delete the user
	err = DeleteUser(conn, client, &user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

}

func TestLogin(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    "testingapilogin@testing.org",
		Password: "PasswordPassword1#",
		Username: "test-user",
	}

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	user.Password = "PasswordPassword1#"

	// login the user
	var token string
	token, err = Login(conn, client, user, "127.0.0.1", "Firefox , Windows 10")

	if err != nil {
		t.Error("The user was not logged in", err)
		return
	}

	if token == "" {
		t.Error("The token is empty")
		return
	}

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

}

func TestLoginWrongPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    "testingapiloginwrongpass@testing.org",
		Password: "PasswordPassword1#",
		Username: "test-user",
	}

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	user.Password = "WrongPassword12#"

	// login the user
	var token string
	token, err = Login(conn, client, user, "127.0.0.1", "Firefox , Windows 10")

	if err == nil {
		t.Error("The user was logged in with wrong password")
		return
	}

	if err.Code != utils.HTTP_STATUS_FORBIDDEN {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	if token != "" {
		t.Error("The token is not empty")
		return
	}

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

}

func TestLoginWrongEmail(t *testing.T) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    "testingapiloginwrongmail@testing.org",
		Password: "PasswordPassword1#",
		Username: "test-user",
	}

	// login the user
	_, err := Login(conn, client, &user, "127.0.0.1", "Firefox , Windows 10")

	if err == nil {
		t.Error("The login was successful")
		return
	}
}
