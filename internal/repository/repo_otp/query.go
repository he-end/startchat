package repootp

const (
	// QueInsertOtpFromEmail = "insert into otp_requests(email, otp_code, expires_at) values($1, $2, now() + interval '5 minutes');"
	// email, otp_code
	queInsertOtpFromEmail = "insert into otp_requests(email, otp_code, expires_at) values($1, $2, $3);"
	// phone, otp_code, expires_at
	// queInsertOtpFromPhone = "insert into otp_requests(phone, otp_code, expires_at) values($1, $2, $3;"
	// email ; *
	queGetOtpFromEmail = "select * from otp_requests where email = $1"
	// phone ; *
	queGetOtpFromPhone = "select * from otp_requests where email = $1"
	// phone
	// queDeleteOtpFromPhone = "delete from otp_requests where phone = $1;"
	// email
	queDeleteOtpFromEmail = "delete from otp_requests where email = $1;"
	// email, otp_code
	queDeleteSpecificOtpCode = "delete from otp_requests where email = $1 and otp_code = $2;"
	// email, otp_code
	queVerifiedOtpFromEmail = "update otp_requests set verified = TRUE where email = $1 and otp_code = $2"
	// email
	// queVerifiedOtpFromPhone = "update otp_requests set verified = TRUE where phone = $1"
)
