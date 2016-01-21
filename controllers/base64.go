package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract"
	"github.com/otiai10/marmoset"
	"github.com/otiai10/ocrserver/config"
)

// Base64 ...
func Base64(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w, true)

	var body = new(struct {
		Base64    string `json:"base64"`
		Trim      string `json:"trim"`
		Languages string `json:"languages"`
		Whitelist string `json:"whitelist"`
	})

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}

	tempfile, err := ioutil.TempFile("", config.AppName()+"-")
	if err != nil {
		render.JSON(http.StatusInternalServerError, err)
		return
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	if len(body.Base64) == 0 {
		render.JSON(http.StatusBadRequest, fmt.Errorf("base64 string required"))
		return
	}
	body.Base64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(body.Base64, "")
	b, err := base64.StdEncoding.DecodeString(body.Base64)
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}
	tempfile.Write(b)

	// TODO: refactor gosseract
	body.Languages = "eng"

	result := gosseract.Must(gosseract.Params{
		Src:       tempfile.Name(),
		Languages: body.Languages,
		Whitelist: body.Whitelist,
	})
	render.JSON(http.StatusOK, map[string]interface{}{
		"result":  strings.Trim(result, body.Trim),
		"version": config.Version(),
	})
}
