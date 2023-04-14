package error

type User int

const (
	USER_ALREADY_EXISTS            = 600
	INVALID_PASSWORD               = 601
	SHORT_PASSWORD                 = 602
	NO_SPECIAL_CHARACTERS_PASSWORD = 603
	NO_MAYUS_MINUS_PASSWORD        = 604
	USER_NOT_UPDATED               = 605
	USER_NOT_FOUND                 = 606
	USER_NOT_DELETED               = 607
  NO_UPPER_LOWER_PASSWORD        = 608
)
