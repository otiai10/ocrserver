package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract"
	"github.com/otiai10/marmoset"
	"github.com/otiai10/ocrserver/config"
)

var (
	imgexp = regexp.MustCompile("^image")
)

// FileUpload ...
func FileUpload(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w, true)

	whitelist := r.FormValue("whitelist")
	trim := r.FormValue("trim")
	// Get uploaded file
	r.ParseMultipartForm(32 << 20)
	// upload, h, err := r.FormFile("file")
	upload, _, err := r.FormFile("file")
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}
	defer upload.Close()

	// Create physical file
	tempfile, err := ioutil.TempFile("", config.AppName()+"-")
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	// Make uploaded physical
	if _, err := io.Copy(tempfile, upload); err != nil {
		render.JSON(http.StatusInternalServerError, err)
		return
	}

	result := gosseract.Must(gosseract.Params{
		Src:       tempfile.Name(),
		Languages: "eng",
		Whitelist: whitelist,
	})

	render.JSON(http.StatusOK, map[string]interface{}{
		"result":  strings.Trim(result, trim),
		"version": config.Version(),
	})
}
