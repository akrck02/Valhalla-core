package utils

import (
	"strings"

	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/lang"
)

const MINIMUM_CHARACTERS_FOR_PASSWORD = 16
const MINIMUM_CHARACTERS_FOR_EMAIL = 5

type validateResult struct {
	Response error.User
	Message  string
}

// Check if the given email is valid
// following the next rules:
//
//		[-] At least 5 characters
//		[-] At least one @
//		[-] At least one .
//
//	 [param] email : string: email to check
//
//	 [return] the email is valid or not
func ValidateEmail(email string) validateResult {

	if len(email) < MINIMUM_CHARACTERS_FOR_EMAIL {
		return validateResult{
			Response: error.SHORT_EMAIL,
			Message:  "Email must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_EMAIL) + " characters",
		}
	}

	if !strings.Contains(email, "@") {
		return validateResult{
			Response: error.NO_AT_EMAIL,
			Message:  "Email must have at least one @",
		}
	}

	if !strings.Contains(email, ".") {
		return validateResult{
			Response: error.NO_DOT_EMAIL,
			Message:  "Email must have at least one .",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check if the given password is valid
// following the next rules:
//
//		[-] At least 16 characters
//		[-] At least one special character
//		[-] At least one number
//
//	 [param] password : string: password to check
//
//	 [return] the password is valid or not
func ValidatePassword(password string) validateResult {

	if len(password) < MINIMUM_CHARACTERS_FOR_PASSWORD {
		return validateResult{
			Response: error.SHORT_PASSWORD,
			Message:  "Password must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_PASSWORD) + " characters",
		}
	}

	if !ContainsSpecialCharacters(password) {
		return validateResult{
			Response: error.NO_SPECIAL_CHARACTERS_PASSWORD,
			Message:  "Password must have at least one special character",
		}
	}

	if IsLowerCase(password) {
		return validateResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one uppercase character",
		}
	}

	if IsUpperCase(password) {
		return validateResult{
			Response: error.NO_UPPER_LOWER_PASSWORD,
			Message:  "Password must have at least one lowercase character",
		}
	}

	if !ContainsNumbers(password) {
		return validateResult{
			Response: error.NO_ALPHANUMERIC_PASSWORD,
			Message:  "Password must have at least one number",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check if the given name string is valid
// following the next rules:
//
//		[-] At least 2 characters
//		[-] At most 50 characters
//
//	 [param] name : string: name to check
//
//	 [return] the name is valid or not
func ValidateName(name string) validateResult {

	if len(name) < 2 {
		return validateResult{
			Response: error.SHORT_NAME,
			Message:  "Name must have at least 2 characters",
		}
	}

	if len(name) > 50 {
		return validateResult{
			Response: error.LONG_NAME,
			Message:  "Name must have at most 50 characters",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check if the given description string is valid
// following the next rules:
//
//		[-] At least 8 characters
//		[-] At most 500 characters
//
//	 [param] description : string: description to check
//
//	 [return] the description is valid or not
func ValidateDescription(description string) validateResult {

	if len(description) < 2 {
		return validateResult{
			Response: error.SHORT_DESCRIPTION,
			Message:  "Description must have at least 2 characters",
		}
	}

	if len(description) > 500 {
		return validateResult{
			Response: error.LONG_DESCRIPTION,
			Message:  "Description must have at most 500 characters",
		}
	}

	return validateResult{
		Response: 200,
		Message:  "Ok.",
	}
}
