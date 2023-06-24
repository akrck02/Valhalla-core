package services

import (
	"github.com/akrck02/valhalla-core/error"
)

type validatorResponse struct {
	Response error.Validator
	Message  string
}
