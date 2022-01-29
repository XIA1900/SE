package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppConfig struct {
	Server struct {
		Port string
	}

	Database struct {
		IP       string
		Port     string
		Username string
		Password string
		Db1      string
		Db2      string
	}

	Redis struct {
		IP       string
		Port     string
		Password string
		DB       int
	}

	JWT struct {
		SecretKey string
		Expires   int
	}

	Logger struct {
		LowestLevel string
		Path        struct {
			Debug string
			Info  string
			Warn  string
			Error string
		}
	}
}

var appConfig AppConfig

func LoadAppConfig() {
	appConfig = AppConfig{}
	appConfigData, _ := ioutil.ReadFile("./config/application.yaml")
	err := yaml.Unmarshal(appConfigData, &appConfig)
	if err != nil {
		panic("Load Application Config Fail")
	}
}

func GetAppConfig() AppConfig {
	return appConfig
}
