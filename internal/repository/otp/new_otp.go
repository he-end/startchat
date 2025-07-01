package otp

import (
	"math/rand"
	"strconv"
)



func generateNewOtp() string {
	max := 999999
	min := 100000
	newNumber := rand.Intn(max-min+1) + min
	otp := strconv.Itoa(newNumber)
	return otp
}
