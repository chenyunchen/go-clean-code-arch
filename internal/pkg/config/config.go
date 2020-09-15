package config

import "time"

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	Context  ContextConfig  `mapstructure:"context"`
}

type ContextConfig struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

type ServerConfig struct {
	Port uint `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
