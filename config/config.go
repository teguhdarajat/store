package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	Database string `mapstructure:"DB_CONN"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := &Config{
		Port:     viper.GetString("PORT"),
		Database: viper.GetString("DB_CONN"),
	}

	return config
}
