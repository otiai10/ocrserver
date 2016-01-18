package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract"
	"github.com/otiai10/ocrserver/config"
)

var (
	imgexp = regexp.MustCompile("^image")
)

// FileUpload ...
func FileUpload(w http.ResponseWriter, r *http.Request) {

	whitelist := r.FormValue("whitelist")
	trim := r.FormValue("trim")
	// Get uploaded file
	r.ParseMultipartForm(32 << 20)
	// upload, h, err := r.FormFile("file")
	upload, _, err := r.FormFile("file")
	if err != nil {
		RenderError(w, http.StatusBadRequest, err)
		return
	}
	defer upload.Close()

	// Validate content type
	/*
		contenttype := h.Header.Get("Content-Type")
		if !imgexp.MatchString(contenttype) {
			RenderError(w, http.StatusBadRequest, fmt.Errorf("invalid content type: %s", contenttype))
			return
		}
	*/

	// Create physical file
	tempfile, err := ioutil.TempFile("", config.AppName()+"-")
	if err != nil {
		RenderError(w, http.StatusBadRequest, err)
		return
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	// Make uploaded physical
	if _, err := io.Copy(tempfile, upload); err != nil {
		RenderError(w, http.StatusInternalServerError, err)
		return
	}

	result := gosseract.Must(gosseract.Params{
		Src:       tempfile.Name(),
		Languages: "eng",
		Whitelist: whitelist,
	})

	Render(w, http.StatusOK, map[string]interface{}{
		"result": strings.Trim(result, trim),
	})
}
