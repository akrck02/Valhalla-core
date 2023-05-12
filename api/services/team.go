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

	err := ownerExists(team.Owner, conn, client)

	if err != nil {
		return err
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)
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

func DeleteTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	/*
		objID, err := utils.StringToObjectId(id)

		teams := client.Database(db.CurrentDatabase).Collection(db.TEAM)
		teams.DeleteMany(conn, bson.M{"_id": team.Owner})
	*/
	return nil
}

func EditTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	objID, err := utils.StringToObjectId(team.ID)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	coll.FindOneAndUpdate(conn, bson.M{"_id": objID}, bson.M{"$set": bson.M{
		"name":        team.Name,
		"description": team.Description,
		"profilepic":  team.ProfilePic,
	},
	})

	return nil
}

func ChangeOwner(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	if utils.IsEmpty(team.Owner) {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_OWNER),
			Message: "Team requires an owner",
		}
	}

	objID, err := utils.StringToObjectId(team.ID)

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	error := ownerExists(team.Owner, conn, client)

	if error != nil {
		return error
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	result := coll.FindOneAndUpdate(conn, bson.M{"_id": objID}, bson.M{"$set": bson.M{
		"owner": team.Owner,
	},
	})

	err = result.Err()

	if err != nil {
		return &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.UPDATE_ERROR),
			Message: "Bad object id",
		}
	}

	return nil
}

func AddMember(conn context.Context, client *mongo.Client, team models.Team) *models.Error {
	return nil
}

func RemoveMember(conn context.Context, client *mongo.Client, team models.Team) *models.Error {
	return nil
}

func GetTeams(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	return nil
}

func GetTeam(conn context.Context, client *mongo.Client, team models.Team) *models.Error {

	return nil
}

func ownerExists(owner string, conn context.Context, client *mongo.Client) *models.Error {

	coll := client.Database(db.CurrentDatabase).Collection(db.USER)
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
