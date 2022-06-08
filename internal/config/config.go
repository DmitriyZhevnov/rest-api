package config

import (
	"sync"

	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool  `yaml:"is_debug"`
	Listen  Listen `yaml:"listen"`
}

type Listen struct {
	BindIP string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
