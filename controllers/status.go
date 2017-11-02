package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
)

const version = "0.2.0"

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w, true).JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello!",
		"version": version,
	})
}
