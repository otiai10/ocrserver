package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/otiai10/marmoset"
	. "github.com/otiai10/mint"
)

func base64server() *httptest.Server {
	r := marmoset.NewRouter()
	r.POST("/", Base64)
	return httptest.NewServer(r)
}

func TestBase64(t *testing.T) {

	s := base64server()

	var body = new(struct {
		Base64    string `json:"base64"`
		Trim      string `json:"trim"`
		Languages string `json:"languages"`
		Whitelist string `json:"whitelist"`
	})

	f, err := os.Open("../test/data/001-base64.txt")
	Expect(t, err).ToBe(nil)
	defer f.Close()
	raw, err := ioutil.ReadAll(f)
	Expect(t, err).ToBe(nil)
	body.Base64 = string(raw)
	body.Trim = "\n"

	b, _ := json.Marshal(body)
	buf := bytes.NewBuffer(b)
	res, err := http.Post(s.URL, "application/json", buf)
	Expect(t, err).ToBe(nil)
	defer res.Body.Close()

	resp := new(struct {
		Result  string `json:"result"`
		Version string `json:"version"`
	})
	err = json.NewDecoder(res.Body).Decode(resp)
	Expect(t, err).ToBe(nil)
	Expect(t, resp.Result).ToBe("ocrserver")
}
