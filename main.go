package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/otiai10/ocrserver/config"
	"github.com/otiai10/ocrserver/controllers"
	"github.com/otiai10/ocrserver/router"
)

func init() {
	configfile := flag.String("conf", "", "config file")
	flag.Parse()
	if *configfile != "" {
		if err := config.InitWithFile(*configfile); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {

	r := router.New()

	// API
	r.GET("/status", controllers.Status)
	r.POST("/base64", controllers.Base64)
	r.POST("/file", controllers.FileUpload)

	r.Static("/assets", config.ProjectPath("assets"))

	// Sample Page
	r.GET("/", controllers.Index)

	logger := &Logger{Debug: config.IsDebug()}

	logger.Printf("[%s]\tListenAndServe\t%s\n", time.Now().Format(time.RFC3339), config.Port())
	http.ListenAndServe(config.Port(), logger.filter(r))
}

// Logger ...
type Logger struct {
	Debug   bool
	handler http.Handler
}

// filter ...
func (logger *Logger) filter(handler http.Handler) *Logger {
	logger.handler = handler
	return logger
}

// Printf ...
func (logger *Logger) Printf(format string, v ...interface{}) {
	if logger.Debug {
		fmt.Printf(format, v...)
	}
}

func (logger *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf("[%s]\t%s\t%s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.String()) // とりあえず
	logger.handler.ServeHTTP(w, r)
}
