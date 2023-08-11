package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/create-room", HandleCreateRoom)
	// router.HandleFunc("/join-room", handleJoinRoom);
	// router.HandleFunc("/send-message", handleSendMessage);

	return router
}
