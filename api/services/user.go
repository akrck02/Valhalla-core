package services

import (
	"context"
	"strings"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/lang"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type validateResult struct {
	Response error.User
	Message  string
}

type EmailChangeRequest struct {
	Email    string `json:"email"`
	NewEmail string `json:"new_email"`
}

const MINIMUM_CHARACTERS_FOR_PASSWORD = 16
const MINIMUM_CHARACTERS_FOR_EMAIL = 5

var SPECIAL_CHARATERS = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", ";", ":", "'", ",", ".", "<", ">", "?", "/", "`", "~"}

// Register user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | *models.User: user to register
//
// [return] *models.Error: error if any
func Register(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	var checkedPass = validatePassword(user.Password)

	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	checkedPass = validateEmail(user.Email)

	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	coll := client.Database(db.DATABASE_NAME).Collection(db.USER)
	found := mailExists(user.Email, conn, coll)

	if found.Email != "" {

		return &models.Error{
			Code:    utils.HTTP_STATUS_CONFLICT,
			Error:   int(error.USER_ALREADY_EXISTS),
			Message: "User already exists",
		}
	}

	user.Password = utils.EncryptSha256(user.Password)

	// register user on database
	_, err := coll.InsertOne(conn, user)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_ALREADY_EXISTS),
			Message: "User already exists",
		}
	}

	return nil
}

// Login user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to login
// [param] ip | string: ip address of the user
// [param] address | string: user agent of the user
//
// [return] string: auth token --> *models.Error: error if any
func Login(conn context.Context, client *mongo.Client, user models.User, ip string, address string) (string, *models.Error) {

	coll := client.Database(db.DATABASE_NAME).Collection(db.USER)
	found := authorizationOk(user.Username, user.Password, conn, coll)

	if found.Email == "" {
		return "", &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Message: "Invalid credentials",
		}
	}

	device := models.Device{Address: ip, UserAgent: address}
	token, err := AddUserDevice(conn, client, found, device)

	if err != nil {
		return "", &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Message: "Cannot generate your auth token",
		}
	}

	return token, nil
}

// Edit user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to edit
//
// [return] *models.Error: error if any
func EditUser(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	users := client.Database(db.DATABASE_NAME).Collection(db.USER)

	// validate email
	if user.Email != "" {
		checkedPass := validateEmail(user.Email)

		if checkedPass.Response != 200 {
			return &models.Error{
				Code:    utils.HTTP_STATUS_FORBIDDEN,
				Error:   int(checkedPass.Response),
				Message: checkedPass.Message,
			}
		}
	}

	// validate password
	if user.Password != "" {
		checkedPass := validatePassword(user.Password)

		if checkedPass.Response != 200 {
			return &models.Error{
				Code:    utils.HTTP_STATUS_FORBIDDEN,
				Error:   int(checkedPass.Response),
				Message: checkedPass.Message,
			}
		}
	}

	// update user on database
	res, err := users.UpdateOne(conn, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"username": user.Username, "password": user.Password}})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User not updated",
		}
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "Users not found",
		}
	}

	return nil
}

// Delete user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to delete
//
// [return] *models.Error: error if any
func DeleteUser(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	users := client.Database(db.DATABASE_NAME).Collection(db.USER)

	// delete user on database
	deleteResult, err := users.DeleteOne(conn, bson.M{"email": user.Email})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_DELETED),
			Message: "User not deleted",
		}
	}

	if deleteResult.DeletedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "User not found",
		}
	}

	return nil
}

// Change password logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to change password
//
// [return] *models.Error: error if any
func ChangeUserPassword(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	users := client.Database(db.DATABASE_NAME).Collection(db.USER)

	var checkedPass = validatePassword(user.Password)
	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	user.Password = utils.EncryptSha256(user.Password)

	// update user on database
	users.UpdateOne(conn, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"password": user.Password}})

	return nil
}

// Change email logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to change email
//
// [return] *models.Error: error if any
func ChangeUserEmail(conn context.Context, client *mongo.Client, mail EmailChangeRequest) *models.Error {

	// update user on database
	users := client.Database(db.DATABASE_NAME).Collection(db.USER)
	var checkedEmail = validateEmail(mail.NewEmail)
	if checkedEmail.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedEmail.Response),
			Message: checkedEmail.Message,
		}

	}

	updateStatus, err := users.UpdateOne(conn, bson.M{"email": mail.Email}, bson.M{"$set": bson.M{"email": mail.NewEmail}})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User not updated" + err.Error(),
		}
	}

	if updateStatus.MatchedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "User not found",
		}
	}

	if updateStatus.ModifiedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User not updated",
		}
	}

	// update user devices on database
	devices := client.Database(db.DATABASE_NAME).Collection(db.DEVICE)

	updateStatus, err = devices.UpdateMany(conn, bson.M{"user": mail.Email}, bson.M{"$set": bson.M{"user": mail.NewEmail}})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User devices not updated",
		}
	}

	if updateStatus.MatchedCount != 0 && updateStatus.ModifiedCount == 0 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_UPDATED),
			Message: "User devices not updated",
		}
	}

	return nil
}

// Check if the given password is valid
// following the next rules:
//
//		[-] At least 16 characters
//		[-] At least one special character
//		[-] At least one number
//
//	 [param] password : string: password to check
//
//	 [return] the password is valid or not
func validatePassword(password string) validateResult {

	if len(password) < MINIMUM_CHARACTERS_FOR_PASSWORD {
		return validateResult{
			Response: error.SHORT_PASSWORD,
			Message:  "Password must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_PASSWORD) + " characters",
		}
	}

	if !utils.ContainsAny(password, SPECIAL_CHARATERS) {
		return validateResult{
			Response: error.NO_SPECIAL_CHARACTERS_PASSWORD,
			Message:  "Password must have at least one special character",
		}
	}

	if utils.IsLowerCase(password) {
		return validateResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one uppercase character",
		}
	}

	if utils.IsUpperCase(password) {
		return validateResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one lowercase character",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check if the given email is valid
// following the next rules:
//
//		[-] At least 5 characters
//		[-] At least one @
//		[-] At least one .
//
//	 [param] email : string: email to check
//
//	 [return] the email is valid or not
func validateEmail(email string) validateResult {

	if len(email) < MINIMUM_CHARACTERS_FOR_EMAIL {
		return validateResult{
			Response: error.SHORT_EMAIL,
			Message:  "Email must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_EMAIL) + " characters",
		}
	}

	if !strings.Contains(email, "@") {
		return validateResult{
			Response: error.NO_AT_EMAIL,
			Message:  "Email must have at least one @",
		}
	}

	if !strings.Contains(email, ".") {
		return validateResult{
			Response: error.NO_DOT_EMAIL,
			Message:  "Email must have at least one .",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check email on database
//
//	[param] email | string The email to check
//	[param] conn | context.Context The connection to the database
//
//	[return] model.User : The user found or empty
func mailExists(email string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "email", Value: email}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result
}

// Get if the given credentials are valid
//
//	[param] username | string : The username to check
//	[param] password | string : The password to check
//	[param] conn | context.Context : The connection to the database
//
//	[return] model.User : The user found or empty
func authorizationOk(username string, password string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "username", Value: username}, {Key: "password", Value: utils.EncryptSha256(password)}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result
}
