package error

type Team int

const (
	NO_PERMISSION       = 630
	TEAM_ALREADY_EXISTS = 631
	EMPTY_TEAM_NAME     = 632
	NO_OWNER            = 633
	OWNER_DOESNT_EXIST  = 634
	BAD_OBJECT_ID       = 635
	UPDATE_ERROR        = 636
	TEAM_NOT_FOUND      = 637
)
