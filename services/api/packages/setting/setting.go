package setting

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
	. "github.com/ymzuiku/hit"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	Port        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	ENV := If(os.Getenv("ENV") == "production", "production", "development")
	pathConfig := fmt.Sprintf("config/%s.ini", ENV)
	cfg, err = ini.Load(pathConfig)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.ini': %v", err)
	}
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
