package mock

func Email() string {
	return "thelegendof@lumberjack.com"
}

func EmailNotDot() string {
	return "thelegendof@lumberjackcom"
}

func EmailNotAt() string {
	return "thelegendoflumberjack.com"
}

func EmailShort() string {
	return "a@s."
}

func Password() string {
	return "PasswordPassword1#"
}

func PasswordShort() string {
	return "Pass1#"
}

func PasswordNotUpperCase() string {
	return "passwordpassword1#"
}

func PasswordNotLowerCase() string {
	return "PASSWORDPASSWORD1#"
}

func PasswordNotNumber() string {
	return "PasswordPassword#"
}

func PasswordNotSpecialChar() string {
	return "PasswordPassword1"
}

func Username() string {
	return "theCrab03"
}
