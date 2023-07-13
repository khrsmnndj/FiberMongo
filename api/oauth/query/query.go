package query

import (
	// "context"
	// "log"
	// "time"
	// "encoding/json"
	"fmt"
	// "reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"

	"github.com/khrsmnndj/FiberMongo/database/mongo/db/connect"
	"github.com/khrsmnndj/FiberMongo/api/oauth/entity"
)

var user *mongo.Collection = connect.MP().Collection("user")

func FindAllUser(c *fiber.Ctx) []entity.User {
	var results []entity.User
	query := c.Queries()
	filter := bson.M{"email": query["email"]}
	cursor, err := user.Find(c.Context(), filter)
	if err != nil {
		fmt.Println(err)
		return results
	}

	if err = cursor.All(c.Context(), &results); err != nil {
	    panic(err)
	}
	if err := cursor.Close(c.Context()); err != nil {
	    panic(err)
	}
	
	return results
}