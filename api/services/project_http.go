package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

func CreateProjectHttp(c *gin.Context) (*models.Response, *models.Error) {

	client := db.CreateClient()
	conn := db.Connect(*client)
	defer db.Disconnect(*client, conn)

	project := &models.Project{}
	err := c.ShouldBindJSON(project)
	if err != nil {
		return nil, &models.Error{
			Status:  utils.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	projectError := CreateProject(conn, client, project)
	if projectError != nil {
		return nil, projectError
	}

	return &models.Response{
		Code: utils.HTTP_STATUS_OK,
		Response: gin.H{
			"message": "Project created",
		},
	}, nil
}
