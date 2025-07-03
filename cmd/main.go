package main

import (
	"fmt"
	"sc/internal/repository/repootp"
)

func main() {
	x := repootp.NewOTP("hendri@email.com")
	fmt.Println(x)
}
