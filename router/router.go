package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/create-room", HandleCreateRoom)
	router.GET("/join-room", HandleJoinRoom)
	router.GET("/exit-room", HandleExitRoom)
	router.GET("/send-message", HandleSendMessage)
	router.GET("/get-room", HandleGetRoom)

	return router
}
