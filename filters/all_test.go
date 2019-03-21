package filters

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/otiai10/marmoset"

	. "github.com/otiai10/mint"
)

func TestLogFilter_ServeHTTP(t *testing.T) {

	router := marmoset.NewRouter()
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	buf := bytes.NewBuffer(nil)
	logger := log.New(buf, "[test] ", 0)
	logfilter := &LogFilter{Logger: logger}
	logfilter.SetNext(router)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	rec := httptest.NewRecorder()

	logfilter.ServeHTTP(rec, req)
	b, _ := ioutil.ReadAll(rec.Body)
	Expect(t, string(b)).ToBe("Hello")
	Expect(t, buf.String()).ToBe("[test] GET /\n")
}
