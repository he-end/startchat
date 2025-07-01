package main

import (
	"fmt"
	"sc/internal/repository/otp"
)

func main() {

	err := otp.SendOTP("asdas")
	if err != nil {
		fmt.Println(err.Error())
	}
	
}
