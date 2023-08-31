package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TanmayPatil105/go-chat/config"
	"github.com/TanmayPatil105/go-chat/database"
	"github.com/TanmayPatil105/go-chat/router"
	"github.com/TanmayPatil105/go-chat/utils"
	"github.com/gin-gonic/gin"
)

var server http.Server

func main() {

	// Connect to Database and Start the Server
	Init()

	// Wait for Terminate signal
	Wait()

	// Disconnect
	Disconnect()

}

func Init() {
	// Get config
	config.GetConfig()

	// Mongo Client
	database.InitDB()

	gin.SetMode(config.AppConfig.GinMode)
	router := router.SetupRouter()

	utils.SetupCronJob()

	server := http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: router,
	}

	go func() {
		fmt.Println("Listening on port", config.AppConfig.Port)
		server.ListenAndServe()
	}()

}

func Wait() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Panicln("Server Shutdown error : ", err)
	}

	log.Println("\nServer Gracefully Stopped!")

	utils.StopCronJob()

	database.DisconnectDB()
}
