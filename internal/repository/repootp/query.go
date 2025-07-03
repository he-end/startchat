package repootp

const (
	// QueInsertOtpFromEmail = "insert into otp_requests(email, otp_code, expires_at) values($1, $2, now() + interval '5 minutes');"
	// email, otp_code
	QueInsertOtpFromEmail = "insert into otp_requests(email, otp_code, expires_at) values($1, $2, $3);"
	// phone, otp_code, expires_at
	QueInsertOtpFromPhone = "insert into otp_requests(phone, otp_code, expires_at) values($1, $2, $3;"
	// email ; *
	QueGetOtpFromEmail = "select * from otp_requests where email = $1"
	// phone ; *
	QueGetOtpFromPhone = "select * from otp_requests where email = $1"
	// phone
	QueDeleteOtpFromPhone = "delete from otp_requests where phone = $1;"
	// email
	QueDeleteOtpFromEmail = "delete from otp_requests where email = $1;"
	// email
	QueVerifiedOtpFromEmail = "update otp_requests set verified = TRUE where email = $1"
	// email
	QueVerifiedOtpFromPhone = "update otp_requests set verified = TRUE where phone = $1"
)
