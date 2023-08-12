package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/TanmayPatil105/go-chat/config"
	"github.com/TanmayPatil105/go-chat/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Room struct {
	SessionId    string    `bson:"sessionId"`
	Owner        string    `bson:"owner, omitempty"`
	Participants []string  `bson:"participants"`
	Messages     []Message `bson:"messages"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func HandleCreateRoom(c *gin.Context) {

	owner := c.Query("owner")
	if owner == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No owner identified",
		})
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("New Session uuid : ", uuid)

	client := database.MongoClient

	options := options.CreateCollection()
	db := client.Database(config.AppConfig.DatabaseName)
	err = db.CreateCollection(context.Background(), uuid.String(), options)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create collection",
		})
		return
	}

	newRoom := Room{
		SessionId:    uuid.String(),
		Owner:        owner,
		Participants: []string{owner},
		Messages:     []Message{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	collection := db.Collection(uuid.String())
	_, err = collection.InsertOne(context.Background(), newRoom)

	if err != nil {
		log.Fatal(err)
	}
	// c.JSON(http.StatusCreated, newRoom)

}

func HandleJoinRoom(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No user detected",
		})
		return
	}

	room := c.Query("room")
	if room == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "room is null",
		})
		return
	}

	client := database.MongoClient
	db := client.Database(config.AppConfig.DatabaseName)

	if exists, _ := database.CollectionExists(db, room); !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Room not found",
		})
	}
	// collection := db.Collection(room)

}
