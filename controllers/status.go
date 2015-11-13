package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/otiai10/ocrserver/config"
)

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Hello!",
		"version": config.Version(),
	})
}
