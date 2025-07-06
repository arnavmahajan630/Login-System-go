package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func ConnectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }
	// small ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }
	log.Println("✅ MongoDB connected")
    return client, nil

}

var Client * mongo.Client

func init() {
	var err error
	Client, err = ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("Mongo connection failed: %v", err))
	}
	if Client == nil {
		panic("Mongo ping failed")
	}
}	

func Opencollection(collection  string) * mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDb client is not initialized. connect db first")
	}
	return Client.Database("userdb").Collection(collection)
}
