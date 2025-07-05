package authpassword

import "regexp"

func IsValidPassword(pwd string) bool {
	if len(pwd) < 12 {
		return false
	}

	// Regex untuk huruf besar (A-Z)
	hasUpperCase, _ := regexp.MatchString(`[A-Z]`, pwd)

	// Regex untuk angka (0-9)
	hasNumber, _ := regexp.MatchString(`[0-9]`, pwd)

	// Regex untuk karakter khusus (misal: /, @, !, #, $, %, &, *, dll.)
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, pwd)

	// Return true hanya jika semua aturan terpenuhi
	return hasUpperCase && hasNumber && hasSpecialChar
}
