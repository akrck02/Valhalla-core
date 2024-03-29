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

type MemberChangeRequest struct {
	Team string `json:"teamid"`
	User string `json:"userid"`
}

// Create team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] user | *models.Team: team to create
//
// [return] error: *models.Error: error if any
func CreateTeam(conn context.Context, client *mongo.Client, team *models.Team) *models.Error {

	// Check if team name is empty
	if utils.IsEmpty(team.Name) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_TEAM_NAME),
			Message: "Team cannot be nameless",
		}
	}

	// Check if team name is valid
	checkedName := utils.ValidateName(team.Name)

	if checkedName.Response != 200 {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedName.Response),
			Message: checkedName.Message,
		}
	}

	// Check if team description is empty
	if utils.IsEmpty(team.Description) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_TEAM_DESCRIPTION),
			Message: "Team cannot be descriptionless",
		}
	}

	// Check if team description is valid
	checkedDescription := utils.ValidateDescription(team.Description)

	if checkedDescription.Response != 200 {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(checkedDescription.Response),
			Message: checkedDescription.Message,
		}
	}

	// Check if owner exists
	err1 := userExists(conn, client, team.Owner)

	if err1 != nil {
		return err1
	}

	// Check if team owner is empty
	if utils.IsEmpty(team.Owner) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_OWNER),
			Message: "Team requires an owner",
		}
	}

	// Check if team already exists
	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	found := teamExists(conn, coll, team)

	if found.Name != "" {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.TEAM_ALREADY_EXISTS),
			Message: "Team already exists with name " + team.Name,
		}
	}

	// Create team
	_, err2 := coll.InsertOne(conn, team)

	if err2 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.TEAM_ALREADY_EXISTS),
			Message: "Team already exists",
		}
	}

	return nil
}

// Delete team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to delete
//
// [return] error: *models.Error: error if any
func DeleteTeam(conn context.Context, client *mongo.Client, team *models.Team) *models.Error {

	// Transform team id to object id
	// also check if team id is valid
	objID, err := utils.StringToObjectId(team.ID)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	// Delete team
	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)
	_, err = coll.DeleteOne(conn, bson.M{"_id": objID})

	// Check if team was deleted
	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.TEAM_NOT_FOUND),
			Message: "Team not found",
		}
	}

	return nil
}

// Edit team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func EditTeam(conn context.Context, client *mongo.Client, team *models.Team) *models.Error {

	// Transform team id to object id
	// also check if team id is valid
	objID, err := utils.StringToObjectId(team.ID)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	update := team.PurgedBson(true)
	_, err = coll.UpdateOne(conn, bson.M{"_id": objID}, bson.M{"$set": update})

	// Check if team was updated
	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.UPDATE_ERROR),
			Message: "Could not update team: " + err.Error(),
		}
	}

	return nil
}

// Edit team owner logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func EditTeamOwner(conn context.Context, client *mongo.Client, team *models.Team) *models.Error {

	// Check if team owner is empty
	if utils.IsEmpty(team.Owner) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_OWNER),
			Message: "Team requires an owner",
		}
	}

	// Transform team id to object id
	// also check if team id is valid
	objID, err1 := utils.StringToObjectId(team.ID)

	if err1 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	// Check if owner exists
	err2 := userExists(conn, client, team.Owner)

	if err2 != nil {
		return err2
	}

	// Update owner
	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	result := coll.FindOneAndUpdate(conn, bson.M{"_id": objID}, bson.M{"$set": bson.M{
		"owner": team.Owner,
	},
	})

	// Check if team was updated
	err3 := result.Err()

	if err3 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.UPDATE_ERROR),
			Message: "Could not change owner",
		}
	}

	return nil
}

// Add member to team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func AddMember(conn context.Context, client *mongo.Client, memberChange *MemberChangeRequest) *models.Error {

	// Check if member is empty
	if utils.IsEmpty(memberChange.User) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_MEMBER),
			Message: "Adding a member requires a member",
		}
	}

	// Check if team is empty
	if utils.IsEmpty(memberChange.Team) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.NO_TEAM),
			Message: "Adding a member requires a team",
		}
	}

	// Check if member exists
	err1 := userExists(conn, client, memberChange.User)

	if err1 != nil {
		return err1
	}

	// Transform team id to object id
	// also check if team id is valid
	objID, err2 := utils.StringToObjectId(memberChange.Team)

	if err2 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	// Check if member is already in team or is owner
	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	err3 := isUserMemberOrOwner(conn, client, memberChange)

	if err3 != nil {
		return err3
	}

	// Add member to team
	_, err4 := coll.UpdateOne(conn, bson.M{"_id": objID}, bson.M{"$push": bson.M{
		"members": memberChange.User,
	},
	})

	if err4 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.UPDATE_ERROR),
			Message: "Could not add member",
		}
	}

	return nil
}

// Remove member from team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func RemoveMember(conn context.Context, client *mongo.Client, member *MemberChangeRequest) *models.Error {
	return nil
}

// Get teams logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func GetTeams(conn context.Context, client *mongo.Client, team *models.Team) *models.Error {

	return nil
}

// Get team logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] team | *models.Team: team to edit
//
// [return] error: *models.Error: error if any
func GetTeam(conn context.Context, client *mongo.Client, team *models.Team) (*models.Team, *models.Error) {

	objID, err1 := utils.StringToObjectId(team.ID)

	if err1 != nil {
		return nil, &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)
	var foundTeam models.Team

	err2 := coll.FindOne(conn, bson.M{"_id": objID}).Decode(&foundTeam)

	if err2 != nil {
		return nil, &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.TEAM_NOT_FOUND),
			Message: "Team not found",
		}
	}

	return &foundTeam, nil
}

// Search teams logic
//
// [param] conn | context.Context: connection to the database
// [param] client | *mongo.Client: client to the database
// [param] searchText | *string: text to search
//
// [return] error: *models.Error: error if any
func SearchTeams(conn context.Context, client *mongo.Client, searchText *string) (*[]models.Team, *models.Error) {

	foundTeams := []models.Team{}

	return &foundTeams, nil

}

func userExists(conn context.Context, client *mongo.Client, user string) *models.Error {

	coll := client.Database(db.CurrentDatabase).Collection(db.USER)
	var foundUser models.User

	objID, err1 := utils.StringToObjectId(user)

	if err1 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.BAD_OBJECT_ID),
			Message: "Bad object id",
		}
	}

	err2 := coll.FindOne(conn, bson.M{"_id": objID}).Decode(&foundUser)

	if err2 != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.OWNER_DOESNT_EXIST),
			Message: "User doesn't exists",
		}
	}

	return nil
}

func teamExists(conn context.Context, coll *mongo.Collection, team *models.Team) models.Team {

	filter := bson.D{
		{Key: "name", Value: team.Name},
		{Key: "owner", Value: team.Owner},
	}
	var result models.Team

	coll.FindOne(conn, filter).Decode(&result)

	return result
}

func isUserMemberOrOwner(conn context.Context, client *mongo.Client, request *MemberChangeRequest) *models.Error {

	filterMember := bson.D{
		{Key: "_id", Value: request.Team},
		{Key: "members", Value: bson.D{{Key: "$all", Value: bson.A{request.User}}}},
	}

	filterOwner := bson.D{
		{Key: "_id", Value: request.Team},
		{Key: "owner", Value: request.User},
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.TEAM)

	var result models.Team

	err := coll.FindOne(conn, filterMember).Decode(&result)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.USER_ALREADY_MEMBER),
			Message: "User is already a member of the team",
		}
	}

	err = coll.FindOne(conn, filterOwner).Decode(&result)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.USER_IS_OWNER),
			Message: "User is owner of the team",
		}
	}

	return nil
}
