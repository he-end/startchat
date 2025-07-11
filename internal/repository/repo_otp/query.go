package repootp

const (
	// QueInsertOtpFromEmail = "insert into otp_requests(email, otp_code, expires_at) values($1, $2, now() + interval '5 minutes');"

	// email, otp_code, purpose, expires_at
	queInsertOtpFromEmail = "insert into otp_requests(email, otp_code, purpose, expires_at) values($1, $2, $3, $4);"
	// phone, otp_code, expires_at
	// queInsertOtpFromPhone = "insert into otp_requests(phone, otp_code, expires_at) values($1, $2, $3;"
	// email, purpose ; *
	queGetOtpFromEmail = "select * from otp_requests where email = $1 and purpose = $2 order by created_at desc limit 1;"
	// phone ; *
	queGetOtpFromPhone = "select * from otp_requests where email = $1"
	// phone
	// queDeleteOtpFromPhone = "delete from otp_requests where phone = $1;"
	// email
	queDeleteOtpFromEmail = "delete from otp_requests where email = $1;"
	// email, otp_code
	queDeleteSpecificOtpCode = "delete from otp_requests where email = $1 and otp_code = $2;"
	// email, otp_code
	// queVerifiedOtpFromEmail = "update otp_requests set verified = TRUE where email = $1 and otp_code = $2"
	// email
	// queVerifiedOtpFromPhone = "update otp_requests set verified = TRUE where phone = $1"
	queExistOtpOrder = "select exists ( select email from otp_requests where email = $1);"
	// count otp request
	queCountOtpRequest = "select count(*) from otp_requests where email = $1 and created_at > NOW() - INTERVAL '20 minute';"
	// email, otp_code
	queUpdateStatusVerifyOtp = "update otp_requests set verified = true, verified_at = now() where id = ( select id from otp_requests where email = $1 and verified = false and expires_at > now() order by created_at desc limit 1 ) and otp_code = $2"
	// UPDATE otp_requests
// SET verified = true, verified_at = NOW()
// WHERE id = (
// SELECT id FROM otp_requests
// WHERE email = $1
// AND verified = false
// AND expires_at > NOW()
// ORDER BY created_at DESC
// LIMIT 1
// )
// AND otp_code = $2;

)
