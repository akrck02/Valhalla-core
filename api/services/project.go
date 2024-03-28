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

func CreateProject(conn context.Context, client *mongo.Client, project *models.Project) *models.Error {

	if utils.IsEmpty(project.Name) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PROJECT_NAME),
			Message: "Project name cannot be empty",
		}
	}

	if utils.IsEmpty(project.Description) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PROJECT_DESCRIPTION),
			Message: "Project description cannot be empty",
		}
	}

	if utils.IsEmpty(project.Owner) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PROJECT_OWNER),
			Message: "Owner cannot be empty",
		}
	}

	coll := client.Database(db.CurrentDatabase).Collection(db.PROJECT)
	found := nameExists(project.Name, conn, coll)

	if found.Name != "" {
		return &models.Error{
			Status:  utils.HTTP_STATUS_CONFLICT,
			Error:   int(error.PROJECT_ALREADY_EXISTS),
			Message: "Project already exists",
		}
	}

	_, err := coll.InsertOne(conn, project)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.PROJECT_ALREADY_EXISTS),
			Message: "Project already exists",
		}
	}

	return nil
}

func EditProject(conn context.Context, client *mongo.Client, project models.Project) *models.Error {

	return nil
}

func DeleteProject(conn context.Context, client *mongo.Client, project models.Project) *models.Error {

	if utils.IsEmpty(project.Name) {
		return &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   int(error.EMPTY_PROJECT_NAME),
			Message: "Project name cannot be empty",
		}
	}

	// delete user devices
	devices := client.Database(db.CurrentDatabase).Collection(db.DEVICE)
	_, err := devices.DeleteMany(conn, bson.M{"project": project.Name})

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.PROJECT_NOT_DELETED),
			Message: "Project not deleted",
		}
	}

	projects := client.Database(db.CurrentDatabase).Collection(db.PROJECT)

	var deleteResult *mongo.DeleteResult
	deleteResult, err = projects.DeleteOne(conn, bson.M{"name": project.Name})

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   int(error.PROJECT_NOT_DELETED),
			Message: "Project not deleted",
		}
	}

	if deleteResult.DeletedCount == 0 {
		return &models.Error{
			Status:  utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.PROJECT_NOT_FOUND),
			Message: "Project not found",
		}
	}

	return nil
}

// Get project logic
func GetProject(conn context.Context, client *mongo.Client, project models.Project, found *models.Project) *models.Error { // get project from database

	projects := client.Database(db.CurrentDatabase).Collection(db.PROJECT)

	err := projects.FindOne(conn, bson.M{"name": project.Name}).Decode(&found)

	if err != nil {
		return &models.Error{
			Status:  utils.HTTP_STATUS_NOT_FOUND,
			Error:   int(error.PROJECT_NOT_FOUND),
			Message: "Project not found",
		}
	}

	found = &models.Project{
		Name:        found.Name,
		Description: found.Description,
	}

	return nil
}

func nameExists(name string, conn context.Context, coll *mongo.Collection) models.Project {
	filter := bson.D{{Key: "name", Value: name}}

	var result models.Project
	coll.FindOne(conn, filter).Decode(&result)

	return result
}
