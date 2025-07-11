package servicependinguser

import (
	"github.com/hend41234/startchat/internal/logger"
	repopendingusers "github.com/hend41234/startchat/internal/repository/repo_pending_users"
	"go.uber.org/zap"
)

func Cleaner() bool {
	// expired pending_users
	_, err := repopendingusers.DeleterPendingUsersExpired()
	if err != nil {
		logger.Error("error run delete pending_users which expired", zap.Error(err))
		return false
	}
	_, err = repopendingusers.DeleterPendingUsersVerified()
	if err != nil {
		logger.Error("error run delete pending_users which verified", zap.Error(err))
		return false
	}
	return true
}
