package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server ServerConfig `json:"server"`
	Specs  SpecsConfig  `json:"specs"`
}

type ServerConfig struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
	IdleTimeout  time.Duration `json:"idleTimeout"`
}

type SpecsConfig struct {
	Sources []SpecSource `json:"sources"`
}

type SpecSource struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	readTimeout := getEnvDuration("READ_TIMEOUT", 10*time.Second)
	writeTimeout := getEnvDuration("WRITE_TIMEOUT", 10*time.Second)
	idleTimeout := getEnvDuration("IDLE_TIMEOUT", 60*time.Second)

	return &Config{
		Server: ServerConfig{
			Port:         port,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
		Specs: SpecsConfig{
			Sources: []SpecSource{
				{
					Name:    "example-api",
					Path:    "./specs/example-api.yaml",
					Version: "v1",
				},
			},
		},
	}
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}
