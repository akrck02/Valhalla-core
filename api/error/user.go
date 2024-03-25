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
	INVALID_EMAIL                  = 609
	SHORT_EMAIL                    = 610
	NO_AT_EMAIL                    = 611
	NO_DOT_EMAIL                   = 612
	NO_ALPHANUMERIC_PASSWORD       = 613
	EMAILS_EQUAL                   = 614
	EMPTY_USERNAME                 = 615
	EMPTY_PASSWORD                 = 616
	EMPTY_EMAIL                    = 617
	INVALID_TOKEN                  = 618
	CANNOT_CREATE_VALIDATION_CODE  = 619
	INVALID_VALIDATION_CODE        = 620
	USER_ALREADY_VALIDATED         = 621
	USER_NOT_AUTHORIZED            = 622
)
