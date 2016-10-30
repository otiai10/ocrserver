package config

import "os"

// values ...
type values struct {
	AppName string `json:"appname"`
	Version string `json:"version"`
	Port    string `json:"port"`
	Debug   bool   `json:"debug"`
}

// default config values
const (
	appname = "ocrserver"
	port    = "8080"
	debug   = true
	version = "0.0.1-default" // change this
)

var v = values{
	AppName: appname,
	Port:    port,
	Debug:   debug,
	Version: version,
}

// Port ...
func Port() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":" + v.Port
}

// Version ...
func Version() string {
	return v.Version
}

// AppName ...
func AppName() string {
	if name := os.Getenv("APP_NAME"); name != "" {
		return name
	}
	return v.AppName
}

// IsDebug ...
func IsDebug() bool {
	return v.Debug
}
