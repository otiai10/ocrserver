package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// values ...
type values struct {
	AppName string `toml:"appname"`
	Version string `toml:"version"`
	Port    string `toml:"port"`
	Debug   bool   `toml:"debug"`
}

// default config values
const (
	appname = "ocrserver"
	port    = "9900"
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
	return v.Port
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

// ProjectPath ...
func ProjectPath(p ...string) string {
	_, currfile, _, _ := runtime.Caller(0)
	return filepath.Join(append([]string{filepath.Dir(filepath.Dir(currfile))}, p...)...)
}
