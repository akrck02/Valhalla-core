package services

import (
	"reflect"

	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

type Endpoint struct {
	Path     string                                                 `json:"path"`
	Method   int                                                    `json:"method"`
	Listener func(models.Request) (*models.Response, *models.Error) `json:"listener"`
	Secured  bool                                                   `json:"secured"`
	Params   interface{}                                            `json:"params"`
	IsForm   bool                                                   `json:"isForm"`
}

var endpoints = []Endpoint{

	// User endpoints
	{"user/register", utils.HTTP_METHOD_PUT, RegisterHttp, false, models.User{}, false},
	// {"user/login", utils.HTTP_METHOD_POST, LoginHttp, false},
	// {"user/edit", utils.HTTP_METHOD_POST, EditUserHttp, true},
	// {"user/edit/email", utils.HTTP_METHOD_POST, EditUserEmailHttp, true},
	// {"user/edit/profilepicture", utils.HTTP_METHOD_POST, EditUserProfilePictureHttp, true},
	// {"user/delete", utils.HTTP_METHOD_DELETE, DeleteUserHttp, true},
	// {"user/get", utils.HTTP_METHOD_GET, GetUserHttp, true},
	// {"user/validate", utils.HTTP_METHOD_GET, ValidateUserHttp, false},

	// // Team endpoints
	// {"team/create", utils.HTTP_METHOD_PUT, CreateTeamHttp, true},
	// {"team/edit", utils.HTTP_METHOD_POST, EditTeamHttp, true},
	// {"team/edit/owner", utils.HTTP_METHOD_POST, EditTeamOwnerHttp, true},
	// {"team/delete", utils.HTTP_METHOD_DELETE, DeleteTeamHttp, true},
	// {"team/get", utils.HTTP_METHOD_GET, GetTeamHttp, true},
	// {"team/add/member", utils.HTTP_METHOD_PUT, AddMemberHttp, true},

	// // Role endpoints
	// {"rol/create", utils.HTTP_METHOD_PUT, CreateRoleHttp, true},
	// {"rol/edit", utils.HTTP_METHOD_POST, EditRoleHttp, true},
	// {"rol/delete", utils.HTTP_METHOD_DELETE, DeleteRoleHttp, true},
	// {"rol/get", utils.HTTP_METHOD_GET, GetRoleHttp, true},

	// Ping
	{"ping", utils.HTTP_METHOD_GET, PingHttp, false, nil, false},
}

func Start() {

	log.Logger.WithDebug()
	log.ShowLogAppTitle()

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(utils.HTTP_STATUS_NOT_FOUND, gin.H{"code": utils.HTTP_STATUS_NOT_FOUND, "message": "Not found"})
	})

	router.Use(CORSMiddleware())
	router.Use(SecurityMiddleware())
	router.Use(PanicManagement())

	registerEndpoints(router)

	log.FormattedInfo("API started on https://${0}:${1}${2}", configuration.Params.Ip, configuration.Params.Port, API_COMPLETE)
	state := router.Run(configuration.Params.Ip + ":" + configuration.Params.Port)
	log.Error(state.Error())

}

func registerEndpoints(router *gin.Engine) {

	for _, endpoint := range endpoints {
		switch endpoint.Method {
		case utils.HTTP_METHOD_GET:
			router.GET(API_COMPLETE+endpoint.Path, APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_POST:
			router.POST(API_COMPLETE+endpoint.Path, APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_PUT:
			router.PUT(API_COMPLETE+endpoint.Path, APIResponseManagement(endpoint))
		case utils.HTTP_METHOD_DELETE:
			router.DELETE(API_COMPLETE+endpoint.Path, APIResponseManagement(endpoint))
		}
	}
}

// Manage errors in a generic way passing the function that will be executed
func APIResponseManagement(endpoint Endpoint) func(c *gin.Context) {

	return func(c *gin.Context) {

		var request models.Request = models.Request{}

		request.Authorization = c.Request.Header.Get("Authorization")
		request.IP = c.ClientIP()
		request.UserAgent = c.Request.Header.Get("User-Agent")

		//if is get
		if c.Request.Method == "GET" {
			request.Params = c.Request.URL.Query()
		}

		//if is post or put or delete
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {

			// if is form
			if endpoint.IsForm {
				// read body
				body, err := utils.MultipartToBytes(c, "file")
				if err != nil {
					c.JSON(utils.HTTP_STATUS_BAD_REQUEST, models.Error{
						Code:    utils.HTTP_STATUS_BAD_REQUEST,
						Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
						Message: "Invalid request",
					})
					return
				}

				request.Body = body
				request.Params = c.Request.PostForm
			} else {

				//create a new instance of the struct
				interfaceType := reflect.TypeOf(endpoint.Params)
				params := reflect.New(interfaceType).Interface()

				// bind params
				err := c.ShouldBind(&params)
				if err != nil {
					c.JSON(utils.HTTP_STATUS_BAD_REQUEST, models.Error{
						Code:    utils.HTTP_STATUS_BAD_REQUEST,
						Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
						Message: "Invalid request" + err.Error(),
					})
					return
				}
			}

		}

		result, error := endpoint.Listener(request)
		if error != nil {
			log.FormattedError("Error ${0} : ${1} in ${2}", string(error.Code), error.Message, c.Request.URL.Path)
			c.JSON(error.Code, error)
			return
		}

		c.JSON(utils.HTTP_STATUS_OK, result)
	}

}

func PanicManagement() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {

			var err = c.Errors[0]

			c.JSON(utils.HTTP_STATUS_INTERNAL_SERVER_ERROR, &models.Error{
				Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
				Error:   error.UNEXPECTED_ERROR,
				Message: err.Error(),
			})
		}
	}
}

func ErrorHandler(c *gin.Context, err gin.Error) {

	httpResponse := gin.H{"code": error.UNEXPECTED_ERROR, "message": err.Error()}
	c.AbortWithStatusJSON(500, httpResponse)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var isRegistered = false

		// Check if endpoint is secured
		for _, endpoint := range endpoints {
			if API_COMPLETE+endpoint.Path == c.Request.URL.Path {
				if !endpoint.Secured {
					log.FormattedInfo("Endpoint ${0} is not secured", endpoint.Path)
					return
				}

				isRegistered = true
			}
		}

		if !isRegistered {
			log.FormattedInfo("Endpoint ${0} is not registered", c.Request.URL.Path)
			return
		}

		log.FormattedInfo("Endpoint ${0} is secured", c.Request.URL.Path)

		// Get token
		token := c.Request.Header.Get("Authorization")

		// Check if token is valid
		if token == "" {
			c.AbortWithStatusJSON(
				utils.HTTP_STATUS_FORBIDDEN,
				gin.H{"code": error.INVALID_TOKEN, "message": "Missing token"},
			)
			return
		}

		// Create a database connection
		var client = db.CreateClient()
		var conn = db.Connect(*client)
		defer db.Disconnect(*client, conn)

		// Check if token is valid
		err := IsTokenValid(client, token)

		if err != nil {
			c.AbortWithStatusJSON(
				err.Code,
				gin.H{"code": err.Error, "message": err.Message},
			)

			return
		}
	}
}
