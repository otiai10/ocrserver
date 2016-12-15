package main

import (
	"fmt"
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

	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", config.AppName()), 0)

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)
	r.POST("/base64", controllers.Base64)
	r.POST("/file", controllers.FileUpload)

	r.Static("/assets", "./assets")

	// Sample Page
	r.GET("/", controllers.Index)

	chain := marmoset.NewFilter(r)
	if os.Getenv("OCRSERVER_LOG_ENABLED") == "1" {
		chain.Add(&filters.LogFilter{Logger: logger})
	}
	server := chain.Server()

	logger.Printf("listening on port %s", config.Port())
	err := http.ListenAndServe(config.Port(), server)
	logger.Println(err)
}
