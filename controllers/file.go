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
)

var (
	imgexp = regexp.MustCompile("^image")
)

// FileUpload ...
func FileUpload(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w, true)

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
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	// Make uploaded physical
	if _, err = io.Copy(tempfile, upload); err != nil {
		render.JSON(http.StatusInternalServerError, err)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(tempfile.Name())
	client.Languages = []string{"eng"}
	if langs := r.FormValue("languages"); langs != "" {
		client.Languages = strings.Split(langs, ",")
	}
	if whitelist := r.FormValue("whitelist"); whitelist != "" {
		client.SetWhitelist(whitelist)
	}

	var out string
	switch r.FormValue("format") {
	case "hocr":
		out, err = client.HOCRText()
		render.EscapeHTML = false
	default:
		out, err = client.Text()
	}
	if err != nil {
		render.JSON(http.StatusBadRequest, err)
		return
	}

	render.JSON(http.StatusOK, map[string]interface{}{
		"result":  strings.Trim(out, r.FormValue("trim")),
		"version": version,
	})
}
