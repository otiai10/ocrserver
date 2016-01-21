package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/ocrserver/config"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w).HTML("index", map[string]interface{}{
		"AppName": config.AppName(),
	})
}
