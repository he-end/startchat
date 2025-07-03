package repootp

import (
	"log"
	"math/rand"
	modelotp "sc/internal/model/otp"
	"sc/internal/repository"
	internalutils "sc/internal/utils"
	"strconv"
	"time"
)

func generateNewOtp() string {
	max := 999999
	min := 100000
	newNumber := rand.Intn(max-min+1) + min
	otp := strconv.Itoa(newNumber)
	return otp
}

func NewOTP(emailOrPhone string) (resOTP modelotp.ResOTP) {
	db := repository.DB

	newOTP := generateNewOtp()

	resOTP.OTPCode = newOTP
	resOTP.ExpiresAt = time.Now().Add(time.Minute * 5)

	if email := internalutils.EmailDetetor(emailOrPhone); email {
		resOTP.Email = emailOrPhone
		tx, err := db.Begin()
		if err != nil {
			log.Fatal("error load database")
		}
		defer tx.Rollback()
		_, err = tx.Exec(QueInsertOtpFromEmail, resOTP.Email, resOTP.OTPCode, resOTP.ExpiresAt)
		if err != nil {
			tx.Rollback()
			log.Fatal(err.Error())
		}
		tx.Commit()
	} else {
		resOTP.Phone = emailOrPhone
	}
	return
}
