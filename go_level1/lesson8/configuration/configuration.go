package configuration

import (
	"flag"
	"net/url"
)

type AppConfig struct {
	port         int
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
}

func GetConfig() AppConfig {
	cfg := AppConfig{}

	flag.IntVar(&cfg.port, "app-port", 0000, "port for app")
	flag.StringVar(&cfg.db_url, "db-url", "http://linkdb.ru", "link for db with params")
	flag.StringVar(&cfg.jaeger_url, "jaeger-url", "http://linkj.ru", "link for jaeger")
	flag.StringVar(&cfg.sentry_url, "sentry-url", "http://links.ru", "link for sentry")
	flag.StringVar(&cfg.kafka_broker, "kafka-broker", "kafka:000", "app name:port")
	flag.StringVar(&cfg.some_app_id, "app-id", "testid", "app id")
	flag.StringVar(&cfg.some_app_key, "app-key", "testkey", "key id")

	flag.Parse()

	UrlValidation(cfg.db_url)
	UrlValidation(cfg.jaeger_url)
	UrlValidation(cfg.sentry_url)

	return cfg
}

func UrlValidation(text_url string) string {
	_, err := url.Parse(text_url)
	if err != nil {
		panic(err)
	} else {
		return text_url
	}
}
