package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w).HTML("index", map[string]interface{}{
		"AppName": "ocrserver",
	})
}
