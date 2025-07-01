package internalutils

import "regexp"

func EmailDetetor(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	var emailRegex = regexp.MustCompile(emailRegexPattern)
	return emailRegex.MatchString(email)
}
