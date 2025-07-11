package servicependinguser

import (
	"database/sql"
	"errors"
	"time"

	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
)

// any return
//
// if the return is true, nil. its mean email not in pending user
//
// when email already to use and expires_at before time now() alias is expired, then return is true and the error is "pending expired". server can does re-registration
//
// email already to use, and not expired, then the return is false, and error is "email already to use". ask client to check inbox for next verify
//
// any error, error may from server or any and return is false ny error
func CheckPendingUser(email string) (bool, error) {
	// found pending_user which using that email
	ok, err := repopendingusers.PendingUserExist(email)
	if ok {
		pu, er := repopendingusers.GetPendingUserWithEmail(email) // pending_users data
		if er != nil {
			return false, er // return any error
		}
		if pu.ExpiresAt.Before(time.Now()) {
			return true, errors.New("pending expired")
		} else {
			return false, errors.New("email is already to use")
		}
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, sql.ErrNoRows // any error
		} else {
			return false, err // any error
		}
	}

	return true, nil //
}
