package conf

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

const (
	ApiKey     = "APIKEY"
	ApiUrl     = "APIURL"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBName     = "DB_NAME"
)

var keys = [...]string{
	ApiKey,
	ApiUrl,
	DBHost,
	DBPort,
	DBUser,
	DBPassword,
	DBName,
}

func GetConfigs() (map[string]string, error) {
	enviroment := make(map[string]string, 0)

	// local
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("error loading enviroment")
		return nil, errors.New("error loading enviroment")
	}

	for _, key := range keys {
		enviroment[key] = os.Getenv(key)
	}

	if len(enviroment) == 0 {
		return nil, errors.New("Error loading env values")
	}
	return enviroment, nil
}
