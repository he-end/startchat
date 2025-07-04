package serviceemail

// func SendEmailWithGmail(to []string, subject string, body string) (ok bool) {
// 	require := utils.Utils.EmailConf

// 	auth := smtp.PlainAuth("", require.Username, require.Password, require.Host)
// 	// set recipient
// 	headerSet := fmt.Sprintf("To: %v\r\nSubject: %v\r\n%v\r\n", to[0], subject, body)
// 	msg := []byte(headerSet)
// 	fmt.Println("process sending")
// 	err := smtp.SendMail(require.Host+":"+require.Port, auth, require.Username, to, msg)
// 	if err != nil {
// 		fmt.Println("error : ", err)
// 		return false
// 	}
// 	fmt.Println("success")
// 	return true

// }
