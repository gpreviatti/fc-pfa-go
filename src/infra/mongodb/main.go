package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func GetConnection(ctx context.Context) *mongo.Client {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING"))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDatabase(client *mongo.Client, name string) *mongo.Database {
	return client.Database(name)
}
