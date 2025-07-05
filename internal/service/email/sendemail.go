package serviceemail

import (
	"fmt"
	"net/smtp"
	"os"
	"sc/internal/logger"

	"github.com/joho/godotenv"
)

func SendEmailWithGmail(to []string, subject string, body string) (ok bool) {
	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.Host)

	// set recipient
	// ========== plain/text
	// headerSet := fmt.Sprintf("To: %v\r\nSubject: %v\r\n%v\r\n", to[0], subject, body)

	// ========== text/html
	headerSet := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		to[0], subject, body,
	)
	msg := []byte(headerSet)
	err := smtp.SendMail(conf.Host+":"+conf.Port, auth, conf.Username, to, msg)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return false
	}
	return true
}

type emailConf struct {
	Username string
	Password string
	Host     string
	Port     string
}

var conf emailConf

func init() {

	if err := godotenv.Load("../configs/.emailconfig"); err != nil {
		logger.Error("error load .emailconf")
		return
	}
	conf.Username = os.Getenv("GMAIL_USERNAME")
	conf.Password = os.Getenv("GMAIL_PASSWORD")
	conf.Host = os.Getenv("GMAIL_HOST")
	conf.Port = os.Getenv("GMAIL_PORT")
}
