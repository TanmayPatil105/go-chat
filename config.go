package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	GinMode string
}

func ReadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	port := os.Getenv("PORT")

	Mode := os.Getenv("GIN_MODE")

	var ginMode string

	if Mode == "DEBUG" {
		ginMode = gin.ReleaseMode
	} else if Mode == "RELEASE" {
		ginMode = gin.DebugMode
	}

	return Config{
		Port:    port,
		GinMode: ginMode,
	}
}
