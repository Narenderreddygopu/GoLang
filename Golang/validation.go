package main
import (
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var PasswardRegex = regexp.MustCompile("^a-zA-Z0-9.!#$%&*.+$")

func IsEmailValid(Emailinp string) bool {
	if len(Emailinp) < 3 && len(Emailinp) > 254 {
		return false
	}

	return emailRegex.MatchString(Emailinp)
}

func IsPasswardValid(Passwordinp string) bool {
	if len(Passwordinp) < 3 && len(Passwordinp) > 254 {
		return false
	}

	return PasswardRegex.MatchString(Passwordinp)
}
