package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hend41234/startchat/internal/handler/http/otp"
	"github.com/hend41234/startchat/internal/handler/http/register"
	"github.com/hend41234/startchat/internal/logger"
	mdwlogger "github.com/hend41234/startchat/internal/middleware/logger"
	mdwratelimiter "github.com/hend41234/startchat/internal/middleware/ratelimiter"
	"github.com/hend41234/startchat/internal/router"
)

func main() {
	logger.Init(logger.Config{
		Environment:     "production",
		LogToConsole:    true,
		LogToFile:       true,
		LogToRemote:     false,
		EnableRolling:   true, // rolling log aktif
		LogFilePath:     "logs/app.log",
		MinimumLogLevel: "debug",
	})
	defer logger.Log.Sync()
	r := router.New()
	apiEndpoint(r)
	wrapped := Middleware(r)

	log.Println("server run in localhost:8080")
	http.ListenAndServe(":8080", wrapped)

}

func apiEndpoint(router *router.Router) {
	router.Handle("POST", "/api/v0.1/register", register.ResgisterHandler)
	router.Handle("POST", "/api/v0.1/verify/otp", otp.VerifyOTPHandler)
}
func Middleware(router http.Handler) http.Handler {
	baseRL := mdwratelimiter.NewRateLimiter(60, time.Minute*1)

	var handler http.Handler = router

	handler = mdwlogger.MiddlewareReqID(handler)      // middleware for autogenerate request_id
	handler = baseRL.MiddelwareBaseRateLimit(handler) // middelware rate limit

	return handler
}
