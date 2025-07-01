package otp

var (
	InsertOtpFromEmail   = "insert into otp_requests(email, otp_code, expires_at) values('%v', '%v', now() + interval '5 minutes');"
	InsertOtpFromPhone   = "insert into otp_requests(phone, otp_code, expires_at) values(%v, %v, now() + interval '5 minutes');"
	DeleteOtpFromPhone   = "delete from otp_requests where phone = '%v';"
	DeleteOtpFromEmail   = "delete from otp_requests where email = '%v';"
	VerifiedOtpFromEmail = "update otp_requests set verified = TRUE where email = '%v'"
	VerifiedOtpFromPhone = "update otp_requests set verified = TRUE where phone = '%v'"
)
