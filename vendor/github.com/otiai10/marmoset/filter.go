package marmoset

import (
	"log"
	"net/http"
	"reflect"
)

// Chain ..
type Chain struct {
	current http.Handler
}

// Filter ...
// Remember "Last added, First called"
type Filter struct {
	// http.Handler
	// Next(next http.Handler) http.Handler
	Next http.Handler
}

// NewFilter ...
func NewFilter(root http.Handler) *Chain {
	return &Chain{
		current: root,
	}
}

// Add ...
func (chain *Chain) Add(filter http.Handler) *Chain {
	v := reflect.ValueOf(filter)
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		// pass
	default:
		log.Fatalf("type `%s` is not addressable", v.Type().String())
		// return chain
	}
	if !v.Elem().FieldByName("Next").CanSet() {
		log.Fatalf("type `%s` must have `Next` field", v.Type().String())
		// return chain
	}
	v.Elem().FieldByName("Next").Set(reflect.ValueOf(chain.current))
	chain.current = filter
	return chain
}

// Server ...
func (chain *Chain) Server() http.Handler {
	return chain.current
}
