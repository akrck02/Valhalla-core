package db

import (
	"context"
	"time"

	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URL = "mongodb://"
const MONGO_USER = "admin"
const MONGO_PASSWORD = "p4ssw0rd"
const MONGO_PORT = "27017"

var DATABASE_NAME = "valhalla"

const TEST_DATABASE_NAME = "valhalla-test"

const USER = "user"
const DEVICE = "device"
const TEAM = "team"
const PROJECT = "project"
const TASK = "task"
const NOTE = "note"
const WIKI = "wiki"
const ROLE = "role"

func CreateClient() *mongo.Client {

	var host = configuration.Params.Mongo

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL + MONGO_USER + ":" + MONGO_PASSWORD + "@" + host + ":" + MONGO_PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}

func Connect(client mongo.Client) context.Context {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.FormattedInfo("Database (${0}) connected on mongodb [${1}:${2}]", DATABASE_NAME, configuration.Params.Mongo, MONGO_PORT)
	return ctx
}

func SetupTest() {

	DATABASE_NAME = TEST_DATABASE_NAME
	var client = CreateClient()
	var ctx = Connect(*client)

	log.Info("Dropping database " + DATABASE_NAME)
	client.Database(DATABASE_NAME).Drop(ctx)
	defer Disconnect(*client, ctx)
}

func Disconnect(client mongo.Client, ctx context.Context) {
	defer client.Disconnect(ctx)
}
