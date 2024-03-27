package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type WeatherConfig struct {
	APIKey  string
	BaseURL string
}

// LoadWeatherConfig - load the API Key and the base URL from environment
func LoadWeatherConfig() WeatherConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("LoadWeatherConfig - godotenv.Load failed to load .env - %v\n", err.Error())
	}

	return WeatherConfig{
		APIKey:  os.Getenv("API_KEY"),
		BaseURL: os.Getenv("BASE_URL"),
	}
}

type LogConfig struct {
	LogLevel string
	AppName  string
}

// LoadLogConfig - load the APP_NAME && LOG_LEVEL to toggle logging verbosity
func LoadLogConfig() LogConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("LoadLog - godotenv.Load failed to load .env - %v\n", err.Error())
	}

	return LogConfig{
		AppName:  os.Getenv("APP_NAME"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
