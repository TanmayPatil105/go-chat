package main

import (
	"fmt"
	"net/http"

	"github.com/TanmayPatil105/go-chat/database"
	"github.com/TanmayPatil105/go-chat/router"
	"github.com/gin-gonic/gin"
)

func main(){
	
	config := ReadConfig();

	database.InitDB()

	gin.SetMode(config.GinMode);
	router := router.SetupRouter();

	server := http.Server{
		Addr: ":"+config.Port,
		Handler: router,
	}

	fmt.Println("Listening on port", config.Port);

	server.ListenAndServe()
	
}
