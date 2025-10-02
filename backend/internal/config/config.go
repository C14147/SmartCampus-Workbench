package config

import "time"

type Config struct {
    Server struct {
        Port         string        `mapstructure:"port"`
        ReadTimeout  time.Duration `mapstructure:"read_timeout"`
        WriteTimeout time.Duration `mapstructure:"write_timeout"`
        IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
    } `mapstructure:"server"`
    Database struct {
        MaxOpenConns    int           `mapstructure:"max_open_conns"`
        MaxIdleConns    int           `mapstructure:"max_idle_conns"`
        ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
    } `mapstructure:"database"`
    Redis struct {
        PoolSize    int `mapstructure:"pool_size"`
        MinIdleConns int `mapstructure:"min_idle_conns"`
    } `mapstructure:"redis"`
}

func DefaultConfig() *Config {
    return &Config{
        Server: struct {
            Port         string
            ReadTimeout  time.Duration
            WriteTimeout time.Duration
            IdleTimeout  time.Duration
        }{
            Port: "8080",
            ReadTimeout: 15 * time.Second,
            WriteTimeout: 15 * time.Second,
            IdleTimeout: 60 * time.Second,
        },
        Database: struct {
            MaxOpenConns    int
            MaxIdleConns    int
            ConnMaxLifetime time.Duration
        }{
            MaxOpenConns: 25,
            MaxIdleConns: 25,
            ConnMaxLifetime: 5 * time.Minute,
        },
        Redis: struct {
            PoolSize    int
            MinIdleConns int
        }{
            PoolSize: 10,
            MinIdleConns: 5,
        },
    }
}
