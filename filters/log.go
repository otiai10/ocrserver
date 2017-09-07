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
func (f *LogFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.Logger.Printf("%s %s", r.Method, r.URL.Path)
	f.Next.ServeHTTP(w, r)
}

// SetNext ...
func (f *LogFilter) SetNext(next http.Handler) {
	f.Next = next
}
