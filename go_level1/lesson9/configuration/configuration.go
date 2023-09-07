package configuration

import (
	"encoding/json"
	"net/url"
	"os"
)

type AppConfig struct {
	Port        int    `json:"port"`
	DbURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func GetConfig() AppConfig {

	config, err := os.ReadFile("configuration.json")
	if err != nil {
		panic(err)
	}

	var myConfig AppConfig

	err = json.Unmarshal(config, &myConfig)
	if err != nil {
		panic(err)
	}

	UrlValidation(myConfig.DbURL)
	UrlValidation(myConfig.JaegerURL)
	UrlValidation(myConfig.SentryURL)

	return myConfig
}

func UrlValidation(text_url string) string {
	_, err := url.Parse(text_url)
	if err != nil {
		panic(err)
	} else {
		return text_url
	}
}
