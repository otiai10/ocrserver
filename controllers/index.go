package controllers

import (
	"net/http"

	"github.com/otiai10/ocrserver/config"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, http.StatusOK, "index.html", map[string]interface{}{
		"AppName": config.AppName(),
	})
}
