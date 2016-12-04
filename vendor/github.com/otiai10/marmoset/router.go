// Package marmoset ,reinventing the wheel
package marmoset

import (
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"runtime"
	"strings"
)

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		routes:   map[string]map[string]http.HandlerFunc{},
		regexps:  map[string]map[*regexp.Regexp]http.HandlerFunc{},
		notfound: http.NotFound,
	}
}

// Router ...
type Router struct {
	static   *static
	routes   map[string]map[string]http.HandlerFunc
	regexps  map[string]map[*regexp.Regexp]http.HandlerFunc
	notfound http.HandlerFunc
}

type static struct {
	Path   string
	Server http.Handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if router.static != nil && strings.HasPrefix(req.URL.Path, router.static.Path) {
		router.static.Server.ServeHTTP(w, req)
		return
	}

	handler := router.findHandler(req)
	if handler == nil {
		router.notfound(w, req)
		return
	}

	handler.ServeHTTP(w, req)
}

// findHandler ...
func (router *Router) findHandler(req *http.Request) http.HandlerFunc {
	if methodes, ok := router.routes[req.Method]; ok {
		if handler, ok := methodes[req.URL.Path]; ok {
			return handler
		}
	}
	if methods, ok := router.regexps[req.Method]; ok {
		for exp, handler := range methods {
			if exp.MatchString(req.URL.Path) {
				matched := exp.FindAllStringSubmatch(req.URL.Path, -1)
				if req.Form == nil {
					req.Form = url.Values{}
				}
				for i, name := range exp.SubexpNames() {
					req.Form.Add(name, matched[0][i])
				}
				return handler
			}
		}
	}
	return nil
}

// add ...
func (router *Router) add(method string, path string, handler http.HandlerFunc) *Router {
	if ok, compiled := isRegexpPath(path); ok {
		return router.addRegexpRoute(method, path, compiled, handler)
	}
	if _, ok := router.routes[method]; !ok {
		router.routes[method] = map[string]http.HandlerFunc{}
	}
	if _, ok := router.routes[method][path]; ok {
		log.Fatalf("route duplicated on `%s %s`", method, path)
	}
	router.routes[method][path] = handler
	return router
}

// isRegexpPath ...
func isRegexpPath(path string) (bool, *regexp.Regexp) {
	compiled, err := regexp.Compile(path)
	if err != nil || compiled == nil {
		return false, nil
	}
	pathParameterExpression := regexp.MustCompile("\\(\\?P\\<[^>]+\\>[^)]+\\)")

	all := strings.Split(path, "/")
	for i, part := range all {
		if pathParameterExpression.MatchString(part) {
			// If the expression is the last part of URL path,
			// The end expression should be appended in Regexp.
			if i == len(all)-1 {
				compiled = regexp.MustCompile(compiled.String() + "$")
			}
			return true, compiled
		}
	}
	return false, nil
}

// addRegexpRoute ...
func (router *Router) addRegexpRoute(method, path string, pathCompiled *regexp.Regexp, handler http.HandlerFunc) *Router {
	if _, ok := router.regexps[method]; !ok {
		router.regexps[method] = map[*regexp.Regexp]http.HandlerFunc{}
	}
	router.regexps[method][pathCompiled] = handler
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

// NotFound ...
func (router *Router) NotFound(handler http.HandlerFunc) *Router {
	router.notfound = handler
	return router
}

// Static ...
func (router *Router) Static(p, dir string) *Router {
	_, f, _, _ := runtime.Caller(1)
	fs := http.FileServer(http.Dir(path.Join(path.Dir(f), dir)))
	router.static = &static{
		Path:   p,
		Server: http.StripPrefix(p, fs),
	}
	return router
}

// StaticRelative ...
func (router *Router) StaticRelative(p string, relativeDir string) *Router {
	// TODO: is this needed?? ;)
	_, f, _, _ := runtime.Caller(1)
	return router.Static(p, path.Join(path.Dir(f), relativeDir))
}
