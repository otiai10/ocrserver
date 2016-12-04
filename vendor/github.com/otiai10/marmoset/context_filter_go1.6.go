// +build appengine

package marmoset

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"
)

var shared = &RequestContextMap{
	contextmap: map[*http.Request]context.Context{},
	locker:     sync.Mutex{},
}

// Context to hide `shared`
func Context() *RequestContextMap {
	return shared
}

// ContextFilter ...
// Only this `ContextFilter` can access `shared` itself.
// Add this filter for the last of your filter chain.
type ContextFilter struct {
	Filter
}

func (f *ContextFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// shared.Set(r, r.Context())
	shared.Set(r, context.Background())
	defer shared.Flush(r)
	f.Next.ServeHTTP(w, r)
}
