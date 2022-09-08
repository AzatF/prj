package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	LogLevel     string `env:"LOG-LEVEL" env-required:"false" env-default:"trace"`
	DataPath     string `env:"DATA-PATH" env-required:"true"`
	BasePath     string `env:"BASE-PATH" env-required:"true"`
	Alpha2       string `env:"ALPHA2" env-required:"true"`
	Host         string `env:"HOST" env-required:"true"`
	Port         string `env:"PORT" env-required:"true"`
	MMSHost      string `env:"MMS-HOST" env-required:"true"`
	MMSPort      string `env:"MMS-PORT" env-required:"true"`
	SupportHost  string `env:"SUPPORT-HOST" env-required:"true"`
	SupportPort  string `env:"SUPPORT-PORT" env-required:"true"`
	IncidentHost string `env:"INCIDENT-HOST" env-required:"true"`
	IncidentPort string `env:"INCIDENT-PORT" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		log.Printf("read application config from path %s", path)

		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
