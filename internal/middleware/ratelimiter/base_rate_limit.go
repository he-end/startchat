package mdwratelimiter

import (
	"net/http"
	"sc/internal/internalutils"
	"sc/internal/logger"
	"time"

	"go.uber.org/zap"
)

func (rl *RateLimiter) MiddelwareBaseRateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.URL.Path, ex)
		// fmt.Println("path with BASE rate limit : ", r.URL.Path)
		// fmt.Println(r.URL.Path, " ", ex)
		host := internalutils.GetClientIP(r)
		if host == "" {
			logger.Error("host not fund or cant find the host")
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		val, exist := rl.Request.Load(host)
		if !exist {
			// fmt.Println("add new clieet : ", host)
			logger.Info("new client", zap.String("IP", host))
			rl.Request.Store(host, 1)
			time.AfterFunc(rl.Interval, func() {
				rl.Request.Delete(host)
			})
		} else {
			count := val.(int)
			if count >= rl.Limit {
				logger.Info("too many request", zap.String("IP", host))
				http.Error(w, "too many request, try again later.", http.StatusTooManyRequests)
				return
			}
			rl.Request.Store(host, count+1)
		}
		next.ServeHTTP(w, r)
	})
}
