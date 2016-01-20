package main

import (
	"log"
	"net/http"
	"os"

	"github.com/otiai10/marmoset"

	"github.com/otiai10/ocrserver/config"
	"github.com/otiai10/ocrserver/controllers"
	"github.com/otiai10/ocrserver/filters"
)

var logger *log.Logger

func main() {

	logger = log.New(os.Stdout, "[ocrserver] ", 0)

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)
	r.POST("/base64", controllers.Base64)
	r.POST("/file", controllers.FileUpload)

	r.Static("/assets", config.ProjectPath("assets"))

	// Sample Page
	r.GET("/", controllers.Index)

	server := marmoset.NewFilter(r).
		Add(&filters.LogFilter{Logger: logger}).
		Server()

	err := http.ListenAndServe(config.Port(), server)
	logger.Println(err)
}
