package marmoset

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var templates *template.Template

// P is just a short hand of `map[string]interface{}`
// I don't want to write `map[string]interface{}` so repeatedly... :(
type P map[string]interface{}

// LoadViews ...
func LoadViews(p string) *template.Template {
	_, f, _, _ := runtime.Caller(1)
	viewpath := path.Join(path.Dir(f), p) + "/"
	exp := regexp.MustCompile("[^/]+\\.html$")

	pool := template.New("")

	filepath.Walk(viewpath, func(fullpath string, info os.FileInfo, err error) error {
		if exp.MatchString(fullpath) {
			name := strings.Replace(strings.Replace(fullpath, viewpath, "", -1), filepath.Ext(fullpath), "", -1)
			tmp, err := template.ParseFiles(fullpath)
			if err != nil {
				panic(err)
			}
			if _, err = pool.AddParseTree(name, tmp.Tree); err != nil {
				panic(err)
			}
		}
		return nil
	})

	templates = pool

	return templates
}
