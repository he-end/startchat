package router

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Route struct {
	method  string
	pattern string
	handler HandlerFunc
}

type Router struct {
	routes []Route
}

func New() *Router {
	return &Router{}
}

func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{method, pattern, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.pattern == req.URL.Path && route.method == req.Method {
			route.handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
