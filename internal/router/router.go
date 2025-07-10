package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hend41234/startchat/internal/internalutils"
	"github.com/hend41234/startchat/internal/logger"

	"go.uber.org/zap"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// base of route
type Route struct {
	method  string
	pattern string
	handler HandlerFunc
}

// list of route / all route will save in here
type Router struct {
	routes []Route
}

func New() *Router {
	return &Router{}
}

// register handler in here
func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{method, pattern, handler})
}

// filtering method every request
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.pattern == req.URL.Path {
			if route.method != req.Method {
				w.Write(resMethodNotAllowed())
				return
			}
			log.Println(req.URL.Path, ":", internalutils.GetClientIP(req))
			route.handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

type methodNotAllowedModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// response method not allowed default
func resMethodNotAllowed() (byteRes []byte) {
	res := methodNotAllowedModel{Code: http.StatusMethodNotAllowed, Message: "method not allowed"}
	byteRes, err := json.Marshal(res)
	if err != nil {
		logger.Error("decode template response error", zap.Error(err))
		return
	}
	return
}
