package repopendingusers

const (
	// email, password, ip_address, token, expires_at
	queAddPendingUser = "insert into pending_users(email, password, ip_address, token, expires_at) values($1, $2, $3, $4, $5)"
	// email, token ; *
	queGetPendingUser = "select * from pendin_users where email = $1 and token = $2"
	// email, token
	queDeletePendingUser = "delete from pending_users where email = $1 and token $2"
)
