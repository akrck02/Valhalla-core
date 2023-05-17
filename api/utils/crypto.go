package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/lang"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/golang-jwt/jwt/v5"
)

// Generate a new auth token
//
// [param] user | models.User | The user
// [param] device | models.Device | The device
//
// [return] string | The token --> error if something went wrong
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

// Encrypt a string using sha256
//
// [param] text | string | The text to encrypt
//
// [return] string | The encrypted text
func EncryptSha256(text string) string {
	plainText := []byte(text)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

// Generate a validation code
//
// [param] text | string | The text to encrypt
//
// [return] string | The encrypted text
func GenerateValidationCode(text string) string {

	// Generate a random string
	randomString, err := GenerateOTP(10)

	if err != nil {
		log.Error(err.Error())
	}

	return randomString
}

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
