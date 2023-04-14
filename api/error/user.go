package error

type User int

const (
	USER_ALREADY_EXISTS            = 600
	INVALID_PASSWORD               = 601
	SHORT_PASSWORD                 = 602
	NO_SPECIAL_CHARACTERS_PASSWORD = 603
	NO_UPPER_LOWER_PASSWORD        = 604
)
