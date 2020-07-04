package DBConnections

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//[Done] This function starts the connection with the mongoBD
func CreateDBConnection(connectionString string) (*mongo.Client, error, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	return client, err, ctx
}
