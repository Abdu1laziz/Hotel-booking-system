package config

import "os"

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBname   string
	}
	User struct {
		Host string
		Port string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Database.User = osGetenv("DB_USER", "postgres")
	c.Database.Password = osGetenv("DB_PASSWORD", "vakhaboff2502")
	c.Database.Host = osGetenv("DB_HOST", "localhost")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "user")

	c.User.Host = osGetenv("USER_HOST", "tcp")
	c.User.Port = osGetenv("USER_PORT", ":8080")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
