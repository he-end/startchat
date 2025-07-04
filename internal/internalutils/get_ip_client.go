package internalutils

import (
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	// try geting from  = X-Forwarded-For
	forwardHost := r.Header.Get("X-Forwarded-For")
	if forwardHost != "" {
		ips := strings.Split(forwardHost, ",")
		return strings.TrimSpace(ips[0])
	}
	// try to get from = X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		// ips := strings.Split(realIP, ",")
		// return strings.TrimSpace(ips[0])
		return realIP
	} else if r.Host != "" {
		host := strings.Split(r.Host, ":")
		return host[0]

	}
	return ""
}
