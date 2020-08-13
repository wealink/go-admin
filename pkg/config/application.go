package config

import (
	"github.com/spf13/viper"
	"time"
)

type Application struct {
	ReadTimeout   time.Duration
	WriterTimeout time.Duration
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
	Domain        string
	IsHttps       bool
	PageSize      int
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:   time.Duration(cfg.GetInt("readTimeout")) * time.Second,
		WriterTimeout: time.Duration(cfg.GetInt("writerTimeout")) * time.Second,
		Host:          cfg.GetString("host"),
		Port:          portDefault(cfg),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		Mode:          cfg.GetString("mode"),
		DemoMsg:       cfg.GetString("demoMsg"),
		Domain:        cfg.GetString("domain"),
		IsHttps:       cfg.GetBool("ishttps"),
		PageSize:      cfg.GetInt("pagesize"),
	}
}

var ApplicationConfig = new(Application)

func portDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8000"
	} else {
		return cfg.GetString("port")
	}
}

func isHttpsDefault(cfg *viper.Viper) bool {
	if cfg.GetString("ishttps") == "" || cfg.GetBool("ishttps") == false {
		return false
	} else {
		return true
	}
}
