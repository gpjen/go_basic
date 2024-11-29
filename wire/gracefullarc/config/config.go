package config

import (
	"fmt"
	"gracefullac/helper"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Name         string
	MaxOpenConns int
	MaxIdleConns int
	MaxLivetime  time.Duration
	TimeZone     string
}

type ServerConfig struct {
	Host string
	Port int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbConfig := DatabaseConfig{
		Host:         helper.GetEnv[string]("DB_HOST", "localhost"),
		Port:         helper.GetEnv[int]("DB_PORT", 3307),
		User:         helper.GetEnv[string]("DB_USER", "root"),
		Password:     helper.GetEnv[string]("DB_PASSWORD", ""),
		Name:         helper.GetEnv[string]("DB_NAME", "test"),
		MaxOpenConns: helper.GetEnv[int]("DB_MAX_OPEN_CONNS", 5),
		MaxIdleConns: helper.GetEnv[int]("DB_MAX_IDLE_CONNS", 5),
		MaxLivetime:  time.Duration(helper.GetEnv[int]("DB_CONN_MAX_LIFETIME", 200)),
		TimeZone:     helper.GetEnv[string]("DB_TIMEZONE", "Asia/Shanghai"),
	}

	serverConfig := ServerConfig{
		Host: helper.GetEnv[string]("SERVER_HOST", "localhost"),
		Port: helper.GetEnv[int]("SERVER_PORT", 8080),
	}

	return &Config{
		Database: dbConfig,
		Server:   serverConfig,
	}, nil
}

func (c *DatabaseConfig) GetDSN(dbType string) string {
	switch dbType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
		)
	case "postgressql":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			c.Host,
			c.User,
			c.Password,
			c.Name,
			c.Port,
			c.TimeZone,
		)
	default:
		return ""
	}
}
