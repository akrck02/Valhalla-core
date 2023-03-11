package services

import (
	"context"

	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUserDevice(conn context.Context, client *mongo.Client, user models.User, device models.Device) (string, error) {

	token, err := utils.GenerateAuthToken(user, device)

	if err != nil {
		return "", err
	}

	coll := client.Database("valhalla").Collection("device")
	device.Token = token
	device.User = user.Email

	found := findDevice(device, conn, coll)

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

func findDevice(device models.Device, conn context.Context, coll *mongo.Collection) models.Device {

	var found models.Device
	coll.FindOne(conn, bson.M{"user": device.User, "address": device.Address, "useragent": device.UserAgent}).Decode(&found)

	return found
}
