package utils

import "regexp"

var isNameValidExp = regexp.MustCompilePOSIX("^[a-zA-Z0-9_-]{4,}$")

func IsValidName(name string) bool {
	isNameValidExp.Longest()

	return name == string(isNameValidExp.Find([]byte(name)))
}
