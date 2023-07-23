package error

type Team int

const (
	NO_PERMISSION          = 630
	TEAM_ALREADY_EXISTS    = 631
	EMPTY_TEAM_NAME        = 632
	NO_OWNER               = 633
	OWNER_DOESNT_EXIST     = 634
	BAD_OBJECT_ID          = 635
	UPDATE_ERROR           = 636
	TEAM_NOT_FOUND         = 637
	SHORT_NAME             = 638
	LONG_NAME              = 639
	SHORT_DESCRIPTION      = 640
	LONG_DESCRIPTION       = 641
	EMPTY_TEAM_DESCRIPTION = 642
	NO_MEMBER              = 643
	NO_TEAM                = 644
	NO_PROJECT             = 645
	USER_IS_OWNER          = 646
	USER_ALREADY_MEMBER    = 647
	TEAM_SEARCH_ERROR      = 648
)
