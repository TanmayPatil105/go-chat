package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	GinMode      string
	DatabaseName string
}

func ReadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	port := os.Getenv("PORT")

	mode := os.Getenv("GIN_MODE")

	database := os.Getenv("DATABASE_NAME")

	var ginMode string

	if mode == "DEBUG" {
		ginMode = gin.ReleaseMode
	} else if mode == "RELEASE" {
		ginMode = gin.DebugMode
	}

	return Config{
		Port:         port,
		GinMode:      ginMode,
		DatabaseName: database,
	}
}
