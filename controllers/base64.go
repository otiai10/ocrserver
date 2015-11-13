package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract"
	"github.com/otiai10/ocrserver/config"
)

// Base64 ...
func Base64(w http.ResponseWriter, r *http.Request) {
	var body = new(struct {
		Base64    string `json:"base64"`
		Trim      string `json:"trim"`
		Languages string `json:"languages"`
		Whitelist string `json:"whitelist"`
	})

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		RenderError(w, http.StatusBadRequest, err)
		return
	}

	tempfile, err := ioutil.TempFile("", config.AppName()+"-")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, err)
		return
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	body.Base64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(body.Base64, "")
	b, _ := base64.StdEncoding.DecodeString(body.Base64)
	tempfile.Write(b)

	// TODO: refactor gosseract
	body.Languages = "eng"

	result := gosseract.Must(gosseract.Params{
		Src:       tempfile.Name(),
		Languages: body.Languages,
		Whitelist: body.Whitelist,
	})
	Render(w, http.StatusOK, map[string]interface{}{
		"result": strings.Trim(result, body.Trim),
	})
}
