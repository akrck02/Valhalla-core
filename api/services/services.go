package services

import (
	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/middleware"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

var ENDPOINTS = []models.Endpoint{

	// User endpoints
	models.EndpointFrom("user/register", utils.HTTP_METHOD_PUT, RegisterHttp, false),
	models.EndpointFrom("user/login", utils.HTTP_METHOD_POST, LoginHttp, false),
	models.EndpointFrom("user/edit", utils.HTTP_METHOD_POST, EditUserHttp, true),
	models.EndpointFrom("user/edit/email", utils.HTTP_METHOD_POST, EditUserEmailHttp, true),
	models.EndpointFrom("user/edit/profilepicture", utils.HTTP_METHOD_POST, EditUserProfilePictureHttp, true),
	models.EndpointFrom("user/delete", utils.HTTP_METHOD_DELETE, DeleteUserHttp, true),
	models.EndpointFrom("user/get", utils.HTTP_METHOD_GET, GetUserHttp, true),
	models.EndpointFrom("user/validate", utils.HTTP_METHOD_GET, ValidateUserHttp, false),

	// Team endpoints
	models.EndpointFrom("team/create", utils.HTTP_METHOD_PUT, CreateTeamHttp, true),
	models.EndpointFrom("team/edit", utils.HTTP_METHOD_POST, EditTeamHttp, true),
	models.EndpointFrom("team/edit/owner", utils.HTTP_METHOD_POST, EditTeamOwnerHttp, true),
	models.EndpointFrom("team/delete", utils.HTTP_METHOD_DELETE, DeleteTeamHttp, true),
	models.EndpointFrom("team/get", utils.HTTP_METHOD_GET, GetTeamHttp, true),
	models.EndpointFrom("team/add/member", utils.HTTP_METHOD_PUT, AddMemberHttp, true),

	// Role endpoints
	models.EndpointFrom("rol/create", utils.HTTP_METHOD_PUT, CreateRoleHttp, true),
	models.EndpointFrom("rol/edit", utils.HTTP_METHOD_POST, EditRoleHttp, true),
	models.EndpointFrom("rol/delete", utils.HTTP_METHOD_DELETE, DeleteRoleHttp, true),
	models.EndpointFrom("rol/get", utils.HTTP_METHOD_GET, GetRoleHttp, true),

	// System endpoints
	models.EndpointFrom("", utils.HTTP_METHOD_GET, ValhallaCoreInfoHttp, false),
}

// Start API
func Start() {

	// set debug or release mode
	if configuration.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
		log.Logger.WithDebug()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.ShowLogAppTitle()
	router := gin.Default()
	router.NoRoute(middleware.NotFound())
	router.Use(middleware.Request())
	router.Use(middleware.CORS())
	router.Use(middleware.Security(ENDPOINTS, API_COMPLETE))
	router.Use(middleware.Panic())

	registerEndpoints(router)

	log.FormattedInfo("API started on https://${0}:${1}${2}", configuration.Params.Ip, configuration.Params.Port, API_COMPLETE)
	state := router.Run(configuration.Params.Ip + ":" + configuration.Params.Port)
	log.Error(state.Error())

}

// Register endpoints
//
// [param] router | *gin.Engine: router
func registerEndpoints(router *gin.Engine) {

	for _, endpoint := range ENDPOINTS {
		switch endpoint.Method {
		case utils.HTTP_METHOD_GET:
			router.GET(API_COMPLETE+endpoint.Path, middleware.APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_POST:
			router.POST(API_COMPLETE+endpoint.Path, middleware.APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_PUT:
			router.PUT(API_COMPLETE+endpoint.Path, middleware.APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_DELETE:
			router.DELETE(API_COMPLETE+endpoint.Path, middleware.APIResponseManagement(endpoint))
		}
	}
}
