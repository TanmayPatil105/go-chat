package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default();

	// router.HandleFunc("/create-room", handleCreateRoom);
	// router.HandleFunc("/join-room", handleJoinRoom);
	// router.HandleFunc("/send-message", handleSendMessage);

	return router;
}
