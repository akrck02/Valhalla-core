package services

import (
	"context"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailChangeRequest struct {
	Email    string `json:"email"`
	NewEmail string `json:"new_email"`
}

// Register user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | *models.User: user to register
//
// [return] *models.Error: error if any
func Register(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	if utils.IsEmpty(user.Email) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_EMAIL),
			Message: "Email cannot be empty",
		}
	}

	if utils.IsEmpty(user.Password) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PASSWORD),
			Message: "Password cannot be empty",
		}
	}

	if utils.IsEmpty(user.Username) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_USERNAME),
			Message: "Username cannot be empty",
		}
	}

	var checkedPass = utils.ValidatePassword(user.Password)

	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	checkedPass = utils.ValidateEmail(user.Email)

	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.USER)
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

	coll := client.Database(db.CurrentDatabase).Collection(db.USER)
	found := authorizationOk(user.Email, user.Password, conn, coll)

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

	if utils.IsEmpty(user.Email) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_EMAIL),
			Message: "Email cannot be empty",
		}
	}

	if utils.IsEmpty(user.Password) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PASSWORD),
			Message: "Password cannot be empty",
		}
	}

	if utils.IsEmpty(user.Username) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_USERNAME),
			Message: "Username cannot be empty",
		}
	}

	users := client.Database(db.CurrentDatabase).Collection(db.USER)

	// validate email
	if user.Email != "" {
		checkedPass := utils.ValidateEmail(user.Email)

		if checkedPass.Response != 200 {
			return &models.Error{
				Code:    utils.HTTP_STATUS_BAD_REQUEST,
				Error:   int(checkedPass.Response),
				Message: checkedPass.Message,
			}
		}
	}

	// validate password
	if user.Password != "" {
		checkedPass := utils.ValidatePassword(user.Password)

		if checkedPass.Response != 200 {
			return &models.Error{
				Code:    utils.HTTP_STATUS_BAD_REQUEST,
				Error:   int(checkedPass.Response),
				Message: checkedPass.Message,
			}
		}
	}

	// update user on database
	res, err := users.UpdateOne(conn, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"username": user.Username, "password": utils.EncryptSha256(user.Password)}})

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

// Change email logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to change email
//
// [return] *models.Error: error if any
func EditUserEmail(conn context.Context, client *mongo.Client, mail EmailChangeRequest) *models.Error {

	if utils.IsEmpty(mail.Email) || utils.IsEmpty(mail.NewEmail) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_EMAIL),
			Message: "Email cannot be empty",
		}
	}

	// Equal emails
	if mail.Email == mail.NewEmail {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMAILS_EQUAL),
			Message: "The new email is the same as the old one",
		}
	}

	// validate email
	var checkedPass = utils.ValidateEmail(mail.Email)
	if checkedPass.Response != 200 {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedPass.Response),
			Message: checkedPass.Message,
		}
	}

	// update user on database
	users := client.Database(db.CurrentDatabase).Collection(db.USER)
	var checkedEmail = utils.ValidateEmail(mail.NewEmail)
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
	devices := client.Database(db.CurrentDatabase).Collection(db.DEVICE)

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

// Delete user logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user to delete
//
// [return] *models.Error: error if any
func DeleteUser(conn context.Context, client *mongo.Client, user models.User) *models.Error {

	if utils.IsEmpty(user.Email) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_EMAIL),
			Message: "Email cannot be empty",
		}
	}

	// delete user devices
	devices := client.Database(db.CurrentDatabase).Collection(db.DEVICE)
	_, err := devices.DeleteMany(conn, bson.M{"user": user.Email})

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.USER_NOT_DELETED),
			Message: "User not deleted",
		}
	}

	// delete user on database
	users := client.Database(db.CurrentDatabase).Collection(db.USER)

	var deleteResult *mongo.DeleteResult
	deleteResult, err = users.DeleteOne(conn, bson.M{"email": user.Email})

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

// Get user logic
func GetUser(conn context.Context, client *mongo.Client, user models.User, found *models.User) *models.Error { // get user from database

	users := client.Database(db.CurrentDatabase).Collection(db.USER)

	err := users.FindOne(conn, bson.M{"email": user.Email}).Decode(&found)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.USER_NOT_FOUND),
			Message: "User not found",
		}
	}

	found = &models.User{
		Email:    found.Email,
		Username: found.Username,
	}

	return nil
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
func authorizationOk(email string, password string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: utils.EncryptSha256(password)}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result
}
