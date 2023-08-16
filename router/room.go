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
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Struct UserId -> Name
type User struct {
	UserId string
	Name   string
}

type Room struct {
	SessionId    string    `bson:"sessionId"`
	Owner        string    `bson:"owner, omitempty"`
	Participants []User    `bson:"participants"`
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

	Owner := User{
		UserId:  xid.New().String(),
		Name: owner,
	}

	newRoom := Room{
		SessionId:    uuid.String(),
		Owner:        owner,
		Participants: []User{Owner},
		Messages:     []Message{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	collection := db.Collection(uuid.String())

	_, err = collection.InsertOne(context.Background(), newRoom)
	if err != nil {
		log.Fatal(err)
	}

	// @Return SessionId and UserId
	c.JSON(http.StatusCreated, gin.H{
		"SessionId": newRoom.SessionId,
		"UserId": Owner.UserId,
	})
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
		return
	}

	newUser := User{
		UserId:  xid.New().String(),
		Name: user,
	}

	collection := db.Collection(room)

	options := options.Update()

	update := bson.M{
		"$push": bson.M{"participants": newUser},
	}

	_, err := collection.UpdateOne(context.Background(), bson.D{}, update, options)
    if err != nil {
        log.Fatal(err)
    }

	// @Return SessionId and UserId
	c.JSON(http.StatusCreated, gin.H{
		"SessionId": room,
		"UserId": newUser.UserId,
	})

	// DEBUGGING

	// Number of documents (1)
	//
	// filter := bson.D{}
	// options := options.Count()
	// collection := db.Collection(room)

	// count, err := collection.CountDocuments(context.Background(), filter, options)
	// if err != nil {
	// 	log.Fatal(err)
	// }

    // Get room based roomid (collection name)
	//
	// filter := bson.D{}
    // options := options.FindOne()
    // var getroom Room
    // err := db.Collection(room).FindOne(context.Background(), filter, options).Decode(&getroom)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// c.JSON(http.StatusCreated, getroom)
}
