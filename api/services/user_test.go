package services

import (
	"fmt"
	"os"
	"path"
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

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

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

func TestRegisterNotEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with no email")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMPTY_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterNotUsername(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with no username")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMPTY_USERNAME {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}
func TestRegisterNotPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with no password")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMPTY_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterNotDotEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.EmailNotDot(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	err := Register(conn, client, user)

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	if err == nil {
		t.Error("The user was registered with an invalid email")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_DOT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterNotAtEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.EmailNotAt(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with an invalid email")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_AT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterShortMail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.EmailShort(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with an invalid email")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.SHORT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterNotSpecialCharactersPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotSpecialChar(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any special character")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_SPECIAL_CHARACTERS_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterNotUpperCaseLoweCasePassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
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

	user = &models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotUpperCase(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err = Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any uppercase letters ")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not registered")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestRegisterShortPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.PasswordShort(),
		Username: mock.Username(),
	}
	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered with short password")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.SHORT_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}
}

func TestRegisterNotNumbersPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.PasswordNotNumber(),
		Username: mock.Username(),
	}
	err := Register(conn, client, user)

	if err == nil {
		t.Error("The user was registered without any number on the password")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_ALPHANUMERIC_PASSWORD {
		t.Error("The error is not the expected" + err.Message)
		return
	}

}

func TestLogin(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

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

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

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

	var user = &models.User{
		Email:    "wrong" + mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.FormattedInfo("Login with email that does not exist ${0}", user.Email)

	// login the user
	_, err = Login(conn, client, &models.User{
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

func TestDeleteUser(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.Info("User registered")

	// login the user
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

	// delete the user
	log.Info("Deleting user")

	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestDeleteUserNoEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Deleting user")

	err := DeleteUser(conn, client, user)

	if err == nil {
		t.Error("The user was deleted")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMPTY_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not deleted")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestDeleteUserNotFound(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Deleting user")

	err := DeleteUser(conn, client, user)

	if err == nil {
		t.Error("The user was deleted")
		return
	}

	if err.Code != utils.HTTP_STATUS_NOT_FOUND {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User not deleted")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = "xXx" + mock.Email()

	var user = &models.User{
		Email:    email,
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

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

	err = EditUserEmail(conn, client, &emailChangeRequest)

	if err != nil {
		t.Error("The user email was not changed", err)
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

func TestEditUserEmailNoEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var emailChangeRequest = EmailChangeRequest{}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMPTY_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserEmailNoDotEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = mock.EmailNotDot()

	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_DOT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)

}

func TestEditUserEmailNoAtEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = mock.EmailNotAt()

	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_AT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserEmailShortEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = mock.EmailShort()

	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.SHORT_EMAIL {
		t.Error("The error is not the expected" + err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserEmailNotFound(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = "xXx" + mock.Email()

	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_NOT_FOUND {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserEmailExists(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()
	var newEmail = mock.Email() + "xXx"

	var user = &models.User{
		Email:    email,
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering original user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The original user was not registered", err)
		return
	}

	log.Info("Original user registered")
	log.Jump()

	// Create a new user with the new email
	newUser := &models.User{
		Email:    newEmail,
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering new user: ${0}", newUser.Email)
	log.FormattedInfo("Password: ${0}", newUser.Password)
	log.FormattedInfo("Username: ${0}", newUser.Username)

	err = Register(conn, client, newUser)

	if err != nil {
		t.Error("The new user was not registered", err)
		return
	}

	log.Info("New user registered")
	log.Jump()

	// Change the email
	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: newEmail,
	}

	log.FormattedInfo("Changing user email to ${0}", newEmail)

	err = EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_CONFLICT || err.Error != error.USER_ALREADY_EXISTS {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Jump()
	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
	log.Jump()

	// Delete the new user
	log.Info("Deleting the original user")

	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
	log.Jump()

	// Delete the new user
	log.Info("Deleting the new user")

	err = DeleteUser(conn, client, newUser)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")

}

func TestEditUserSameEmail(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email = mock.Email()

	var emailChangeRequest = EmailChangeRequest{
		Email:    email,
		NewEmail: email,
	}

	log.Info("Changing user email")

	err := EditUserEmail(conn, client, &emailChangeRequest)

	if err == nil {
		t.Error("The user email was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.EMAILS_EQUAL {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User email not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserPassword(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password
	err = EditUser(conn, client, user)

	if err != nil {
		t.Error("The user password was not changed: " + err.Message)
		return
	}

	log.Info("User password changed")

	// check if the user can login with the new password
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

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestEditUserPasswordUserNotFound(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.Info("Changing user password")

	err := EditUser(conn, client, user)

	if err == nil {
		t.Error("The user password was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_NOT_FOUND {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User password not changed")
	log.FormattedInfo("Error: ${0}", err.Message)
}

func TestEditUserPasswordShort(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password
	user.Password = mock.PasswordShort()

	err = EditUser(conn, client, user)

	if err == nil {
		t.Error("The user password was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.SHORT_PASSWORD {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User password not changed")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestEditUserPasswordNoLowercase(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password

	user.Password = mock.PasswordNotLowerCase()

	err = EditUser(conn, client, user)

	if err == nil {
		t.Error("The user password was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User password not changed")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestEditUserPasswordNoUppercase(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password

	user.Password = mock.PasswordNotUpperCase()

	err = EditUser(conn, client, user)

	if err == nil {
		t.Error("The user password was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_UPPER_LOWER_PASSWORD {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User password not changed")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestEditUserPasswordNoNumber(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// change the user password

	user.Password = mock.PasswordNotNumber()

	err = EditUser(conn, client, user)

	if err == nil {
		t.Error("The user password was changed")
		return
	}

	if err.Code != utils.HTTP_STATUS_BAD_REQUEST || err.Error != error.NO_ALPHANUMERIC_PASSWORD {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.Info("User password not changed")
	log.FormattedInfo("Error: ${0}", err.Message)

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestEditProfilePicture(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	profilePic, readErr := utils.ReadFile(mock.ProfilePicture())

	if readErr != nil {
		t.Error("The file was not read", readErr)
		return
	}

	err = EditUserProfilePicture(conn, client, user, profilePic)

	if err != nil {
		t.Error("The profile picture was not changed", err)
		return
	}

	log.Info("Profile picture changed")

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestOsDirName(t *testing.T) {
	filepath, err := os.Getwd()
	if err != nil {
		log.Info(err.Error())
	}

	filepath = path.Dir(path.Dir(filepath))

	fmt.Println(filepath)
}

func TestTokenValidation(t *testing.T) {

	// Create a new user
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user = &models.User{
		Username: mock.Username(),
		Email:    mock.Email(),
		Password: mock.Password(),
	}

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// Login the user
	var token string
	token, err = Login(conn, client, user, mock.Ip(), mock.Platform())

	if err != nil {
		t.Error("The user was not logged in", err)
		return
	}

	_, err = middleware.IsTokenValid(client, token)

	if err != nil {
		t.Error("The token was not validated", err)
		return
	}

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

	log.Info("User deleted")
}

func TestTokenValidationInvalidToken(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create a fake token
	token := mock.Token()
	_, err := middleware.IsTokenValid(client, token)

	if err == nil {
		t.Error("The token was validated")
		return
	}

	if err.Code != utils.HTTP_STATUS_FORBIDDEN || err.Error != error.INVALID_TOKEN {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.FormattedInfo("Token not validated, error { http: ${0}, internal: ${1}, message: \"${2}\" }", utils.Int2String(err.Code), utils.Int2String(err.Error), err.Message)

}

func TestTokenValidationInvalidTokenFormat(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create a fake token
	token := mock.Username()
	_, err := middleware.IsTokenValid(client, token)

	if err == nil {
		t.Error("The token was validated")
		return
	}

	if err.Code != utils.HTTP_STATUS_FORBIDDEN || err.Error != error.INVALID_TOKEN {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.FormattedInfo("Token not validated, error { http: ${0}, internal: ${1}, message: \"${2}\" }", utils.Int2String(err.Code), utils.Int2String(err.Error), err.Message)

}

func TestTokenValidationEmptyToken(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create a fake token
	token := ""
	_, err := middleware.IsTokenValid(client, token)

	if err == nil {
		t.Error("The token was validated")
		return
	}

	if err.Code != utils.HTTP_STATUS_FORBIDDEN || err.Error != error.INVALID_TOKEN {
		t.Error("The error is not the expected", err.Message)
		return
	}

	log.FormattedInfo("Token not validated, error { http: ${0}, internal: ${1}, message: \"${2}\" }", utils.Int2String(err.Code), utils.Int2String(err.Error), err.Message)

}

func TestValidationCode(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Create a new user
	var user = &models.User{
		Username: mock.Username(),
		Email:    mock.Email(),
		Password: mock.Password(),
	}

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	// get the user
	user, err = GetUser(conn, client, user, true)

	// validate the user
	err = ValidateUser(conn, client, user.ValidationCode)

	if err != nil {
		t.Error("The user was not validated", err)
		return
	}

	// delete the user
	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted", err)
		return
	}

}
