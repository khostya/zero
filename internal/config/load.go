package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		Env  string `env-required:"true" yaml:"env"`
		HTTP HTTP   `mapstructure:"http"`
		PG   PG     `mapstructure:"postgres"`
	}

	HTTP struct {
		Port         int           `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
		IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	}

	PG struct {
		URL             string
		MaxOpenConns    int           `mapstructure:"max_open_conns"`
		MaxIdleConns    int           `mapstructure:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	}
)

const (
	configPath = "./config/config.yml"
)

func NewConfig() (Config, error) {
	path, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		path = configPath
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	viper.SetConfigType("yml")
	err = viper.ReadConfig(file)
	if err != nil {
		return Config{}, err
	}

	var configuration Config
	err = viper.Unmarshal(&configuration)

	pgURL := os.Getenv("PG_URL")
	if pgURL == "" {
		return Config{}, errors.New("PG_URL environment variable not set")
	}

	configuration.PG.URL = pgURL

	return configuration, err
}

func MustNewConfig() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
