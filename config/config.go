package config

import (
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

// values ...
type values struct {
	AppName string `toml:"appname"`
	Version string `toml:"version"`
	Port    string `toml:"port"`
	Debug   bool   `toml:"debug"`
}

var v = values{
	AppName: "ocrserver",
	Version: "0.0.1-default",
	Port:    ":9900",
	Debug:   true,
}

// InitWithFile ...
func InitWithFile(fpath string) error {
	_, err := toml.DecodeFile(fpath, &v)
	if err != nil {
		return err
	}
	// log.Println(meta)
	return nil
}

// Port ...
func Port() string {
	return v.Port
}

// Version ...
func Version() string {
	return v.Version
}

// AppName ...
func AppName() string {
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
