package controllers

import (
	"encoding/json"
	"net/http"
)

// Render ...
func Render(w http.ResponseWriter, code int, v interface{}) {
	b, _ := json.MarshalIndent(v, "", "\t")
	w.Write(b)
}

// RenderError ...
func RenderError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	b, _ := json.MarshalIndent(map[string]interface{}{
		"error": err.Error(),
	}, "", "\t")
	w.Write(b)
}
