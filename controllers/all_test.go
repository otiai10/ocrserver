package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/otiai10/marmoset"
	. "github.com/otiai10/mint"
)

func testserver() *httptest.Server {
	r := marmoset.NewRouter()
	r.POST("/base64", Base64)
	r.POST("/file", FileUpload)
	r.GET("/status", Status)
	r.GET("/", Index)
	return httptest.NewServer(r)
}

func TestBase64(t *testing.T) {

	s := testserver()

	type Response struct {
		Result  string `json:"result"`
		Version string `json:"version"`
	}

	type RequestBodyForBase64 struct {
		Base64    string `json:"base64"`
		Trim      string `json:"trim"`
		Languages string `json:"languages"`
		Whitelist string `json:"whitelist"`
	}
	body := new(RequestBodyForBase64)

	f, err := os.Open("../test/data/001-base64.txt")
	Expect(t, err).ToBe(nil)
	defer f.Close()
	raw, err := ioutil.ReadAll(f)
	Expect(t, err).ToBe(nil)
	body.Base64 = string(raw)
	body.Trim = "\n"

	b, _ := json.Marshal(body)
	buf := bytes.NewBuffer(b)
	res, err := http.Post(s.URL+"/base64", "application/json", buf)
	Expect(t, err).ToBe(nil)
	defer res.Body.Close()

	resp := new(struct {
		Result  string `json:"result"`
		Version string `json:"version"`
	})
	err = json.NewDecoder(res.Body).Decode(resp)
	Expect(t, err).ToBe(nil)
	Expect(t, resp.Result).ToBe("ocrserver")

	When(t, "no request body provided", func(t *testing.T) {
		res, err := http.Post(s.URL+"/base64", "application/json", nil)
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).Not().ToBe(http.StatusOK)
	})

	When(t, "no base64 string provided", func(t *testing.T) {
		body := new(RequestBodyForBase64)
		b, _ := json.Marshal(body)
		res, err := http.Post(s.URL+"/base64", "application/json", bytes.NewBuffer(b))
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).Not().ToBe(http.StatusOK)
	})

	When(t, "invalid base64 string provided", func(t *testing.T) {
		body := new(RequestBodyForBase64)
		body.Base64 = "_____invalid_____"
		b, _ := json.Marshal(body)
		res, err := http.Post(s.URL+"/base64", "application/json", bytes.NewBuffer(b))
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).Not().ToBe(http.StatusOK)
	})

	When(t, "languages and whitelist provided", func(t *testing.T) {
		body := new(RequestBodyForBase64)
		body.Languages = "eng,jpn"
		body.Whitelist = "012345ver"
		body.Base64 = string(raw)
		b, _ := json.Marshal(body)
		res, err := http.Post(s.URL+"/base64", "application/json", bytes.NewBuffer(b))
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).ToBe(http.StatusOK)
		defer res.Body.Close()
		resp := new(Response)
		json.NewDecoder(res.Body).Decode(resp)
		Expect(t, resp.Result).ToBe("00r5erver")
	})
}

func TestFileUpload(t *testing.T) {

	s := testserver()

	type Response struct {
		Result  string `json:"result"`
		Version string `json:"version"`
	}

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	src, _ := os.Open("../test/data/ocrserver.png")
	part, _ := w.CreateFormFile("file", "ocrserver.png")
	io.Copy(part, src)
	w.Close()
	res, err := http.Post(s.URL+"/file", w.FormDataContentType(), body)
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusOK)

	When(t, "multipart/form-data file not provided", func(t *testing.T) {
		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		w.Close()
		res, err := http.Post(s.URL+"/file", w.FormDataContentType(), body)
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).ToBe(http.StatusBadRequest)
	})

	When(t, "hocr, langs, and whitelist provided", func(t *testing.T) {
		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		src, _ := os.Open("../test/data/ocrserver.png")
		part, _ := w.CreateFormFile("file", "ocrserver.png")
		io.Copy(part, src)
		w.WriteField("languages", "eng,jpn")
		w.WriteField("whitelist", "ocrserver")
		w.WriteField("format", "hocr")
		w.Close()
		res, err := http.Post(s.URL+"/file", w.FormDataContentType(), body)
		Expect(t, err).ToBe(nil)
		Expect(t, res.StatusCode).ToBe(http.StatusOK)
	})

}

func TestStatus(t *testing.T) {
	s := testserver()
	res, err := http.Get(s.URL + "/status")
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusOK)
}

func TestIndex(t *testing.T) {
	s := testserver()
	res, err := http.Get(s.URL + "/")
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusOK)
}
