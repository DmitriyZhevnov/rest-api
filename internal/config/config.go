package config

import (
	"sync"

	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug *bool `mapstructure:"is_debug"`
	Listen  Listen
	Storage Storage
	Auth    AuthConfig
}

type AuthConfig struct {
	PasswordSalt string
}

type Storage struct {
	MongoDB MongoDB `json:"mongodb"`
}
type MongoDB struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

type Listen struct {
	BindIP string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")

		if err := setUpViper(); err != nil {
			logger.Fatal(err)
		}

		instance = &Config{}
		if err := unmarshal(instance); err != nil {
			logger.Fatal(err)
		}

		if err := fromEnv(instance); err != nil {
			logger.Fatal(err)
		}
	})
	return instance
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("listen", &cfg.Listen); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("stotage.mongodb", &cfg.Storage.MongoDB); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg.Auth.PasswordSalt = viper.GetString("PASSWORD_SALT")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
