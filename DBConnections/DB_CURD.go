package DBConnections

import (
	. "../config"
	. "../schema"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

//[Done] This function inserts a new field in a collection
/*func DbInsert(data interface{}, collection string) bool {
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = mongoClient.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}*/

func BindMap(c *gin.Context) map[string]interface{}{
	m := make(map[string]interface{})
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad binding"})
		os.Exit(1)
	}
	return m
}

func InterfaceToUser(c *gin.Context,userData interface{}) User{
	var user User
	tmpData,err := json.Marshal(userData)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad user data"})
		os.Exit(1)
	}
	err=json.Unmarshal(tmpData,&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad unmarching"})
		os.Exit(1)
	}
	return user
}

func DbInsert(c *gin.Context)  {
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}

	postContentMap := BindMap(c)
	var data User
	data=InterfaceToUser(c,postContentMap["user"])
	collection := postContentMap["collection"].(string)

	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = mongoClient.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "insertion error"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "insertion completed successfully"})
	return
}

//[Done] This function reads one or more value in a field using any other value
//func DbRead(findByKey string, findByValue interface{}, collection string, readKey ...string) (interface{}, error) {
func DbRead(c *gin.Context)  {
	postContentMap:=BindMap(c)
	fmt.Println("3aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println(postContentMap)
	findByKey := postContentMap["findByKey"].(string)
	findByValue := postContentMap["findByValue"].(string)
	collection := postContentMap["collection"].(string)
	readKey := postContentMap["readKey"].([]string)

	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)

	var projection bson.D
	for _, projKey := range readKey {
		projection = append(projection, bson.E{projKey, 1})
	}

	result, err := mongoClient.Find(context.Background(), bson.D{{findByKey, findByValue}}, options.Find().SetProjection(projection))

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"message": "Did not find matching data"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "mongodb find error"})
	}

	for result.Next(context.Background()) {
		var usr User
		// Decode the document
		if err := result.Decode(&usr); err != nil {
			log.Println("cursor.Decode ERROR:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Decode error"})
		}

			c.JSON(http.StatusAccepted, gin.H{"user": usr})
	}
	c.JSON(http.StatusNotFound, gin.H{"user": ""})
}

//[Done?] This function deletes a field in a collection by it's ID
//to be tested
/*func DbDelete(ID primitive.ObjectID, collection string) bool {
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)

	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = mongoClient.DeleteOne(ctx, bson.M{"_id": ID})

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}*/

//[Done] This function updates a certain object inside a field in a collection using ID
/*func DbUpdate(prmUserID primitive.ObjectID, collection string, key string, data interface{}) bool {

	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"_id", prmUserID}}
	update := bson.D{{"$set", bson.D{{key, data}}}}
	result, err := mongoClient.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	Println(result)
	return true
}
*/