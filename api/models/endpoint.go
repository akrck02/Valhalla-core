package models

import "github.com/gin-gonic/gin"

type Endpoint struct {
	Path     string           `json:"path"`
	Method   int              `json:"method"`
	Listener EndpointListener `json:"listener"`
	Secured  bool             `json:"secured"`
}

type EndpointListener func(*gin.Context) (*Response, *Error)

func EndpointFrom(path string, method int, listener EndpointListener, secured bool) Endpoint {
	return Endpoint{
		Path:     path,
		Method:   method,
		Listener: listener,
		Secured:  secured,
	}
}
