package filters

import (
	"log"
	"net/http"
)

// LogFilter ...
type LogFilter struct {
	Logger *log.Logger
	Next   http.Handler
}

// ServeHTTP ...
func (h LogFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("%s %s", r.Method, r.URL.Path)
	h.Next.ServeHTTP(w, r)
}
