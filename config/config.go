package config

import (
	"gopkg.in/gcfg.v1"
	"log"
)

type Config struct {
	Mongo struct {
		Hosts        string
		AuthDatabase string
		AuthUsername string
		AuthPassword string
		Database     string
	}
	Host struct {
		Name string
		Port string
		Path string
	}
	ProductHunt struct {
		Endpoint  string
		ApiKey    string
		ApiSecret string
	}
	CookieStore struct {
		AuthenticationKey string
	}
}

var cfg Config
var loaded bool

func Get() Config {
	if !loaded {
		Load()
	}
	return cfg
}

func Load() {
	err := gcfg.ReadFileInto(&cfg, "/etc/packhunter/packhunter.gcfg")
	if err != nil {
		log.Printf("Failed to read config: %v", err)
	} else {
		loaded = true
	}
}
