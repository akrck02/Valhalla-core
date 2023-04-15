package services

import (
	"context"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddUserDevice adds a new device to the database
// or updates the token if the device already exists
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | models.User: user that owns the device
// [param] device | models.Device: device to add
//
// [return] string: token of the device --> error : The error that occurred
func AddUserDevice(conn context.Context, client *mongo.Client, user models.User, device models.Device) (string, error) {

	token, err := utils.GenerateAuthToken(user, device)

	if err != nil {
		return "", err
	}

	coll := client.Database(db.DATABASE_NAME).Collection(db.DEVICE)
	device.Token = token
	device.User = user.Email

	found := findDevice(conn, coll, device)

	if found.Address != "" {

		log.Debug("Device already exists, updating token")
		coll.ReplaceOne(conn, found, device)

		return token, nil
	}

	log.Debug("Creating new device...")

	_, err = coll.InsertOne(conn, device)

	if err != nil {
		return "", err
	}

	return token, nil
}

// findDevice finds a device in the database
//
// [param] conn | context.Context: connection to the database
// [param] coll | *mongo.Collection: collection to search
// [param] device | models.Device: device to find
//
// [return] models.Device: device found --> error : The error that occurred
func findDevice(conn context.Context, coll *mongo.Collection, device models.Device) models.Device {

	var found models.Device
	coll.FindOne(conn, bson.M{"user": device.User, "address": device.Address, "useragent": device.UserAgent}).Decode(&found)

	return found
}
