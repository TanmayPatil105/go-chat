package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/TanmayPatil105/go-chat/config"
	"github.com/TanmayPatil105/go-chat/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	UserId    string    `bson:"userId"`
	Text      string    `bson:"text"`
	TimeStamp time.Time `bson:"sent_at"`
}

func HandleSendMessage(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
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

	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No message sent",
		})
		return
	}

	timestamp := time.Now()

	message := Message{
		UserId:    userId,
		Text:      text,
		TimeStamp: timestamp,
	}

	client := database.MongoClient
	db := client.Database(config.AppConfig.DatabaseName)

	if exists, _ := database.CollectionExists(db, room); !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Room not found",
		})
		return
	}

	collection := db.Collection(room)

	options := options.Update()

	// Insert message
	update := bson.M{
		"$push": bson.M{"messages": message},
	}

	_, err := collection.UpdateOne(context.Background(), bson.D{}, update, options)
	if err != nil {
		log.Fatal(err)
	}

	// Update updated_at
	update = bson.M{
		"$set": bson.M{"updated_at": timestamp},
	}

	_, err = collection.UpdateOne(context.Background(), bson.D{}, update, options)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, message)
}
