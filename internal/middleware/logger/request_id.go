package mdwlogger

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const RequestIDKey contextKey = "request_id"

func MiddlewareReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper untuk logger
func GetRequestID(ctx context.Context) string {
	reqID, ok := ctx.Value(RequestIDKey).(string)
	if !ok || reqID == "" {
		return "-"
	}
	return reqID
}
