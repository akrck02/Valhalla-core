package services

import (
	"runtime"
	"time"

	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

type ValhallaCoreInfo struct {
	Version    string   `json:"version"`
	License    string   `json:"license"`
	Authors    []string `json:"authors"`
	Copyleft   string   `json:"copyleft"`
	Repository string   `json:"repository"`
	GoVersion  string   `json:"go-version"`
}

func ValhallaCoreInfoHttp(c *gin.Context) (*models.Response, *models.Error) {

	// get go version
	goVersion := runtime.Version()

	return &models.Response{
		Code: utils.HTTP_STATUS_OK,
		Response: ValhallaCoreInfo{
			Version: "1.0.0",
			License: "GNU GPLv3",
			Authors: []string{
				"akrck02",
				"AlejandroMacazaga",
				"AnderRod01",
			},
			Copyleft:   time.Now().Format("2006"),
			Repository: "https://github.com/akrck02/valhalla-core",
			GoVersion:  goVersion,
		},
	}, nil
}
