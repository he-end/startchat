package main

import (
	"sc/internal/logger"
	serviceotp "sc/internal/service/otp"
)

func main() {
	logger.Init(logger.Config{
		Environment:     "production",
		LogToConsole:    true,
		LogToFile:       true,
		LogToRemote:     false,
		EnableRolling:   true, // rolling log aktif
		LogFilePath:     "logs/app.log",
		MinimumLogLevel: "debug",
	})
	defer logger.Log.Sync()

	// mail := "test123@test.com"
	// fmt.Println("=> generate new otp")
	// newOtp, _ := repootp.NewOTP(mail)
	// if len(newOtp.Email) < 1 {
	// 	log.Fatal()
	// }
	// fmt.Println("=> wait for a second")
	// time.Sleep(time.Second * 2)
	// fmt.Println("=> test matching")
	// ok, err := serviceotp.VerifyOtp(mail, newOtp.OTPCode)
	// if !ok {
	// 	if err != nil {
	// 		logger.Error("error mathing otp", zap.Error(err))
	// 		return
	// 	}
	// 	logger.Error("otp not match", zap.String("otp", newOtp.OTPCode))
	// 	return
	// }
	// fmt.Println("otp match")

	serviceotp.SendOTPWithGmail("1010101", "hendri41234@gmail.com")
}
