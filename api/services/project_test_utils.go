package services

import (
	"context"
	"testing"

	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateMockTestProject creates a project for testing purposes
//
// [param] t | *testing.T: Testing object
// [param] conn | context.Context: Database connection
// [param] client | *mongo.Client: Database client
//
// [return] models.Project: Created project
func CreateMockTestProject(t *testing.T, conn context.Context, client *mongo.Client) *models.Project {
	user := RegisterMockTestUser(t, conn, client)

	project := &models.Project{
		Name:        "Test Project",
		Description: "Test Description",
		Owner:       user.Email,
	}

	log.FormattedInfo("Creating project: %v", project.Name)
	err := CreateProject(conn, client, project)

	if err != nil {
		t.Errorf("Error creating project: %v", err)
	}

	t.Log("Project created successfully")
	return project
}

// CreateMockTestProject creates a project for testing purposes
//
// [param] t | *testing.T: Testing object
// [param] conn | context.Context: Database connection
// [param] client | *mongo.Client: Database client
//
// [return] models.Project: Created project
func CreateMockTestProjectWithUser(t *testing.T, conn context.Context, client *mongo.Client, user *models.User) *models.Project {

	project := &models.Project{
		Name:        "Test Project",
		Description: "Test Description",
		Owner:       user.Email,
	}

	log.FormattedInfo("Creating project: %v", project.Name)
	err := CreateProject(conn, client, project)

	if err != nil {
		t.Errorf("Error creating project: %v", err)
	}

	t.Log("Project created successfully")
	return project
}

// CreateTestProjectWithoutOwner creates a project without an owner for testing purposes
//
// [param] t | *testing.T: Testing object
// [param] conn | context.Context: Database connection
// [param] client | *mongo.Client: Database client
func CreateTestProjectWithError(t *testing.T, conn context.Context, client *mongo.Client, project *models.Project, status int, errorcode int) {

	log.FormattedInfo("Creating project: ${0}", project.Name)
	err := CreateProject(conn, client, project)

	if err == nil {
		t.Error("Project created successfully")
		return
	}

	if err.Status != status || err.Error != errorcode {
		t.Errorf("Error code mismatch: %v", err)
		return
	}

	t.Log("Project creation failed as expected")
	log.FormattedInfo("Error creating project: ${0}", err.Message)
}
