package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

type Config struct {
	Environment string

	ServerHost string
	HTTPPort   string

	PostgresHost          string
	PostgresUser          string
	PostgresDatabase      string
	PostgresPassword      string
	PostgresPort          int
	PostgresMaxConnection int32

	DefaultOffset int
	DefaultLimit  int
	SecretKey     string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "localhost:"))
	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", "8000"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "doniy"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "app"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1901dony2003"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	cfg.PostgresMaxConnection = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))
	cfg.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "SECRET"))
	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
