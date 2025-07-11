package repopendingusers

const (
	// email, password, ip_address, token, expires_at
	queAddPendingUser = "insert into pending_users(email, password, token, expires_at) values($1, $2, $3, $4)"
	// email, token ; *
	// queGetPendingUser = "select * from pendin_users where email = $1 and token = $2"
	// email, token ; *
	// queGetPendingUser2 = "select * from pending_users where ip_address = $1 and token = $2"
	// token ; *
	queGetPendingUser3 = "select * from pending_users where token = $1;"
	// email
	queGetPendingUserWithemail = "select * from pending_users where email = $1;"
	// [*] cleaner according expires pending_users
	queCleanerPendingUserExpired = "delete from pending_users where verified = false and expires_at < now();"
	// [*] cleaner according verified pending_users
	queCleanerPendingUserVerified = "delete from pending_users where verified = true and expires_at < now();"
	// exist pending user
	queExistPendigUser = "select exists ( select email from pending_users where email = $1)"
	// token
	queUpdateStatusPending = "update pending_users set verified = true where token = $1;"
	// newPassword, expires_at, token, email
	queUpdateOldRecord = "update pending_users set password = $1, created_at = NOW(), expires_at = $2, token = $3 where email = $4  "
)
