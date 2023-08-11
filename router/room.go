package router

import (
	"context"

	"github.com/TanmayPatil105/go-chat/config"
	"github.com/TanmayPatil105/go-chat/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleCreateRoom(c *gin.Context) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	config := config.ReadConfig()

	// fmt.Println("New Session uuid : ", uuid)

	client := database.MongoClient

	db := client.Database(config.DatabaseName)
	options := options.CreateCollection()
	db.CreateCollection(context.Background(), uuid.String(), options)

	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "Failed to create collection",
		// })
		return
	}

}
