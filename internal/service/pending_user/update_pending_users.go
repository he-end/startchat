package servicependinguser

import repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"



func UpdateStatusPendingUsers(token string)(bool, error){
	ok, err := repopendingusers.UpdateStatusPending(token)
	if err != nil{
		return ok, err
	}
	return true, nil
}