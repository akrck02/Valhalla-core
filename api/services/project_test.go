package services

import (
	"testing"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/mock"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
)

func TestCreateProject(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	CreateMockTestProject(t, conn, client)
}

func TestCreateProjectWithoutOwner(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := models.Project{
		Name:        "Test Project",
		Description: "Test Description",
	}

	CreateTestProjectWithError(t, conn, client, &project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_OWNER)
}

func TestCreateProjectWithoutName(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := &models.Project{
		Description: mock.ProjectDescription(),
		Owner:       mock.Email(),
	}

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_NAME)
}
func TestCreateProjectWithoutDescription(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := &models.Project{
		Name:  mock.ProjectName(),
		Owner: mock.Email(),
	}

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_BAD_REQUEST, error.EMPTY_PROJECT_DESCRIPTION)
}

func TestCreateProjectThatAlreadyExists(t *testing.T) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	user := RegisterMockTestUser(t, conn, client)
	project := CreateMockTestProjectWithUser(t, conn, client, user)

	CreateTestProjectWithError(t, conn, client, project, utils.HTTP_STATUS_CONFLICT, error.PROJECT_ALREADY_EXISTS)
}
