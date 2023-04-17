package utils

import (
	"testing"

	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/mock"
)

func TestValidateCode(t *testing.T) {

	email := mock.Email()
	code := GenerateValidationCode(email)
	log.Info(code)

}
