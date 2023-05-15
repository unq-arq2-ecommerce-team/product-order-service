package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Environment   string          `required:"true" default:"development"`
	Port          int             `required:"true" default:"8080"`
	LogLevel      string          `split_words:"true" default:"DEBUG"`
	MongoURI      string          `split_words:"true" required:"true"`
	MongoDatabase string          `split_words:"true" required:"true"`
	MongoTimeout  time.Duration   `split_words:"true" required:"true"`
	Seller        SellerEndpoints `split_words:"true" required:"true"`
}

type SellerEndpoints struct {
	UrlFindById string `split_words:"true" required:"true"`
}

func LoadConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
