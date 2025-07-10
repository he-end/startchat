package repopendingusers

const (
	// email, password, ip_address, token, expires_at
	queAddPendingUser = "insert into pending_users(email, password, ip_address, token, expires_at) values($1, $2, $3, $4, $5)"
	// email, token ; *
	queGetPendingUser = "select * from pendin_users where email = $1 and token = $2"
	// email, token ; *
	queGetPendingUser2 = "select * from pending_users where ip_address = $1 and token = $2"
	// email, token ; *
	queGetPendingUser3 = "select * from pending_users where token = $1;"
	// email, token
	queDeletePendingUser = "delete from pending_users where email = $1 and token $2"
	// email, token
	queDeletePendingUser2 = "delete from pending_users where email = $1"
	// exist pending user
	queExistPendigUser = "select exists ( select email from pending_users where email = $1)"
)
