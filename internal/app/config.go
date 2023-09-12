package app

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"ENV"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBDsn      string `mapstructure:"DATABASE_URL"`
}

func LoadConfigFromENV() (*Config, error) {
	config := &Config{
		Env:        os.Getenv("ENV"),
		DBDsn:      os.Getenv("DATABASE_URL"),
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
	return config, nil
}

func LoadConfig(path string) (config *Config, err error) {
	env := os.Getenv("ENV")
	viper.AddConfigPath(path)
	viper.SetConfigName("app-" + env + ".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
