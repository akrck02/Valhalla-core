package utils

import (
	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/lang"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAuthToken(user models.User, device models.Device) (string, error) {

	now := getCurrentMillis()

	log.Jump()
	log.Debug("device: " + device.UserAgent + "-" + device.Address)
	log.Debug("username: " + user.Username)
	log.Debug("email: " + user.Email)
	log.Debug("timestamp: " + lang.Int642String(now))
	log.Jump()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"device":    device.UserAgent + "-" + device.Address,
		"username":  user.Username,
		"email":     user.Email,
		"timestamp": now,
	})

	tokenString, err := token.SignedString([]byte(configuration.Params.Secret))
	return tokenString, err
}
