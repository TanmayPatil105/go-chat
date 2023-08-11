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

	"github.com/TanmayPatil105/go-chat/database"
	"github.com/TanmayPatil105/go-chat/router"
	"github.com/gin-gonic/gin"
)

func main() {

	config := ReadConfig()

	// Mongo Client
	client, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	gin.SetMode(config.GinMode)
	router := router.SetupRouter()

	server := http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	go func() {
		fmt.Println("Listening on port", config.Port)
		server.ListenAndServe()
	}()

	// Graceful Server Shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Panicln("Server Shutdown error : ", err)
	}

	log.Println("\nServer Gracefully Stopped!")

	database.DisconnectDB(client)
}
