package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `yaml:"is_debug" env:"Bot_IsDebug" env-default:"false" env-required:"true"`
	IsDevelopment bool `yaml:"is_development" env:"Bot_IsDevelopment" env-default:"false" env-required:"true"`
	Telegram      struct {
		Token string `yaml:"token" env:"Bot_Telegram_Token" env-required:"true"`
	}
	RabbitMQ struct {
		Host     string `yaml:"host" env:"Bot_RabbitMQ_Host" env-required:"true"`
		Port     string `yaml:"port" env:"Bot_RabbitMQ_Port" env-required:"true"`
		Username string `yaml:"username" env:"Bot_RabbitMQ_Username" env-required:"true"`
		Password string `yaml:"password" env:"Bot_RabbitMQ_Password" env-required:"true"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
