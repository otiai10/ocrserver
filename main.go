package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/otiai10/ocrserver/config"
	"github.com/otiai10/ocrserver/controllers"
	"github.com/otiai10/ocrserver/router"
)

func init() {
	configfile := flag.String("conf", "", "Config file")
	flag.Parse()
	if *configfile != "" {
		if err := config.InitWithFile(*configfile); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {

	r := router.New()
	r.GET("/status", controllers.Status)
	r.POST("/base64", controllers.Base64)

	http.ListenAndServe(config.Port(), r)

}
