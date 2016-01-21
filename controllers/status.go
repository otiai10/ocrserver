package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/ocrserver/config"
)

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w, true).JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello!",
		"version": config.Version(),
	})
}
