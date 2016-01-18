package controllers

import "net/http"

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, http.StatusOK, "index.html", nil)
}
