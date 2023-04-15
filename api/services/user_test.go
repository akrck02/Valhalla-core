package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/mock"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
)

func TestRegister(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.Info("User registered")

	// delete the user
	log.Info("Deleting user: " + user.Email)
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")

}

func TestRegisterNotDotEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.EmailNotDot(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	err := Register(conn, client, user)

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	if err == nil {
		t.Error("The user was registered with an invalid email")
		return
	}

	if err.Error != error.NO_DOT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterMinimumCharacters(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.PasswordShort(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with less than minimum characters")
		return
	}

	if err.Error != error.SHORT_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterSpecialCharacters(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotSpecialChar(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any special character")
		return
	}

	if err.Error != error.NO_SPECIAL_CHARACTERS_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterUpperCaseLoweCase(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotLowerCase(),
		Username: mock.Username(),
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

	user = models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotUpperCase(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err = Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestChangeUserPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password
	err = ChangeUserPassword(conn, client, user)

	if err != nil {
		t.Error("The user password was not changed")
		return
	}

	log.Info("User password changed")

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestLogin(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	user.Password = "PasswordPassword1#"
	log.FormattedInfo("New password : ${0}", user.Password)

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

	log.Info("User logged in")
	log.FormattedInfo("Token: ${0}", token)

	// delete the user
	log.Info("Deleting user")
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")

}

func TestLoginWrongPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.Info("User registered")

	user.Password = mock.PasswordShort()

	log.Info("Login with wrong password")
	log.FormattedInfo("Password: ${0}", user.Password)

	// login the user
	var token string
	token, err = Login(conn, client, user, mock.Ip(), mock.Platform())

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

	log.Info("User not logged in")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	log.Info("Deleting user")
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestLoginWrongEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = models.User{
		Email:    "wrong" + mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.FormattedInfo("Login with email that does not exist ${0}", user.Email)

	// login the user
	_, err = Login(conn, client, models.User{
		Email:    mock.Email(),
		Password: user.Password,
	}, mock.Ip(), mock.Platform())

	if err == nil {
		t.Error("The login was successful")
		return
	}

	if err.Code != utils.HTTP_STATUS_FORBIDDEN {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not logged in")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	log.Info("Deleting user")
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

// TODO: add tests for the following functions
// func ChangePassword
// func DeleteUser
// func EditUser

func TestChangeUserEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = "xXx" + mock.Email()

	var user = models.User{
		Email:    email,
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Registering user: " + user.Email)
	log.Info("Password: " + user.Password)
	log.Info("Username: " + user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// login the user to create a device
	var token string
	token, err = Login(conn, client, user, mock.Ip(), mock.Platform())

	if err != nil {
		t.Error("The user was not logged in", err)
		return
	}

	if token == "" {
		t.Error("The token is empty")
		return
	}

	log.Info("User logged in")

	// Change the user email
	log.Info("Changing user email")
	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	err = ChangeUserEmail(conn, client, emailChangeRequest)

	if err != nil {
		t.Error("The user email was not changed ", err)
		return
	}

	log.Info("User email changed")

	user.Email = newEmail

	// delete the user
	log.Info("Deleting user")
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}
