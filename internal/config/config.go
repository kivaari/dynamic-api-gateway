package config

import (
	"github.com/kivaari/dynamic-api-gateway/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Discovery DiscoveryConfig `mapstructure:"discovery"`
	Security  SecurityConfig  `mapstructure:"security"`
}

type ServerConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	GracefulTimeout string `mapstructure:"graceful_timeout"`
}

type DiscoveryConfig struct {
	Type   string       `mapstructure:"type"`
	Consul ConsulConfig `mapstructure:"consul"`
	Static StaticConfig `mapstructure:"static"`
}

type ConsulConfig struct {
	Address string `mapstructure:"address"`
	Prefix  string `mapstructure:"prefix"`
}

type StaticConfig struct {
	Services []ServiceConfig `mapstructure:"services"`
}

type ServiceConfig struct {
	Name   string `mapstructure:"name"`
	Host   string `mapstructure:"host"`
	Target string `mapstructure:"target"`
}

type SecurityConfig struct {
	JWT       JWTConfig       `mapstructure:"jwt"`
	CORS      CORSConfig      `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
}

type JWTConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Secret      string `mapstructure:"secret"`
	TokenHeader string `mapstructure:"token_header"`
}

type CORSConfig struct {
	Enabled        bool     `mapstructure:"enabled"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

type RateLimitConfig struct {
	Enabled           bool `mapstructure:"enabled"`
	RequestsPerSecond int  `mapstructure:"requests_per_second"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Log.Errorf("Unable to decode config: %v", err)
		return nil, err
	}

	return &cfg, nil
}
