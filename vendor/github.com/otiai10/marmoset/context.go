// +build !appengine

package marmoset

import (
	"context"
	"net/http"
	"sync"
)

// RequestContextMap ...
type RequestContextMap struct {
	contextmap map[*http.Request]context.Context
	locker     sync.Mutex
}

// Set ...
func (rctxmap *RequestContextMap) Set(req *http.Request, ctx context.Context) {
	// TODO: Avoid duplicated call of this `if`
	if rctxmap.contextmap == nil {
		rctxmap.contextmap = map[*http.Request]context.Context{}
	}
	rctxmap.locker.Lock()
	defer rctxmap.locker.Unlock()
	rctxmap.contextmap[req] = ctx
}

// Get ...
func (rctxmap *RequestContextMap) Get(req *http.Request) context.Context {
	rctxmap.locker.Lock()
	defer rctxmap.locker.Unlock()
	return rctxmap.contextmap[req]
}

// Flush ...
func (rctxmap *RequestContextMap) Flush(req *http.Request) {
	if _, ok := rctxmap.contextmap[req]; ok {
		rctxmap.locker.Lock()
		delete(rctxmap.contextmap, req)
		rctxmap.locker.Unlock()
	}
}
