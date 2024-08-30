package config

import (
	"os"
)

type Config struct {
	User struct {
		Host string
		Port string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.User.Host = osGetenv("HOST", "localhost")
	c.User.Port = osGetenv("PORT", "8085")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
	