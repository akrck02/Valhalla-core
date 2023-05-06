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

func CreateTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	if utils.IsEmpty(team.Name) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_TEAM_NAME),
			Message: "Team cannot be nameless",
		}
	}

	if utils.IsEmpty(team.Owner) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_OWNER),
			Message: "Team requires an owner",
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)
	err := ownerExists(team.Owner, conn, coll)

	if err != nil {
		return err
	}

	_, err2 := coll.InsertOne(conn, team)

	if err2 != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.TEAM_ALREADY_EXISTS),
			Message: "Team already exists",
		}
	}

	return nil
}

func DeleteTeam(conn context.Context, client *mongo.Client, id string) *models.Error {
	/*
		objID, err := utils.StringToObjectId(id)

		teams := client.Database(db.CurrentDatabase).Collection(db.TEAM)
		teams.DeleteMany(conn, bson.M{"_id": team.Owner})
	*/
	return nil
}

func EditTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	return nil
}

func GetTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	return nil
}

func ownerExists(owner string, conn context.Context, coll *mongo.Collection) *models.Error {

	var foundUser models.User

	objID, err := utils.StringToObjectId(owner)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	err = coll.FindOne(conn, bson.M{"_id": objID}).Decode(&foundUser)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.OWNER_DOESNT_EXIST),
			Message: "Owner doesn't exists",
		}
	}

	return nil
}
