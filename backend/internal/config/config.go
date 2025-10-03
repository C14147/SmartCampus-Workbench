package config

import (
    "github.com/spf13/viper"
)

type ServerConfig struct {
    Port string `mapstructure:"port"`
}

type Config struct {
    Server ServerConfig `mapstructure:"server"`
    JWT    struct {
        Secret string `mapstructure:"secret"`
    } `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
    v := viper.New()
    v.SetConfigName("config")
    v.AddConfigPath(".")
    v.AddConfigPath("./config")
    v.SetDefault("server.port", "8080")
    v.SetDefault("jwt.secret", "change-me-in-production")

    if err := v.ReadInConfig(); err != nil {
        // it's OK if no config file; we'll use defaults and env
    }

    v.AutomaticEnv()

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
