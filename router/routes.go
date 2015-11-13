// Package router ,reinventing the wheel
package router

import (
	"log"
	"net/http"
)

// New ...
func New() *Router {
	return &Router{map[string]map[string]http.HandlerFunc{}}
}

// Router ...
type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodes, ok := router.routes[r.Method]
	if !ok {
		http.NotFound(w, r)
		return
	}
	handler, ok := methodes[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	handler.ServeHTTP(w, r)
}

// add ...
func (router *Router) add(method string, path string, handler http.HandlerFunc) *Router {
	if _, ok := router.routes[method]; !ok {
		router.routes[method] = map[string]http.HandlerFunc{}
	}
	if _, ok := router.routes[method][path]; ok {
		log.Fatalf("route duplicated on `%s %s`", method, path)
	}
	router.routes[method][path] = handler
	return router
}

// GET ...
func (router *Router) GET(path string, handler http.HandlerFunc) *Router {
	return router.add("GET", path, handler)
}

// POST ...
func (router *Router) POST(path string, handler http.HandlerFunc) *Router {
	return router.add("POST", path, handler)
}
