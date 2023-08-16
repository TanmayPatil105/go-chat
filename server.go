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
	"github.com/gin-gonic/gin"
)

func main() {

	config.GetConfig()

	// Mongo Client
	database.InitDB()

	gin.SetMode(config.AppConfig.GinMode)
	router := router.SetupRouter()

	server := http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: router,
	}

	go func() {
		fmt.Println("Listening on port", config.AppConfig.Port)
		server.ListenAndServe()
	}()

	// Graceful Server Shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Panicln("Server Shutdown error : ", err)
	}

	log.Println("\nServer Gracefully Stopped!")

	database.DisconnectDB()
}
