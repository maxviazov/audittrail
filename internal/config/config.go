package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

type logConfig struct {
	ConsoleLevel string `mapstructure:"console_level"`
	FileLevel    string `mapstructure:"file_level"`
}

type Config struct {
	Log logConfig `mapstructure:"logging"`
}

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("logging.console_level", "info")
	v.SetDefault("logging.file_level", "debug")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}
	return &cfg, nil
}
