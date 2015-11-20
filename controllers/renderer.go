package controllers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/otiai10/ocrserver/config"
)

var _templates *template.Template

func init() {
	tpl, err := template.ParseGlob(config.ProjectPath("views/*.html"))
	if err != nil {
		panic(err)
	}
	_templates = tpl
}

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

// RenderHTML ...
func RenderHTML(w http.ResponseWriter, code int, name string, data interface{}) {
	w.WriteHeader(code)
	_templates.ExecuteTemplate(w, name, data)
}
