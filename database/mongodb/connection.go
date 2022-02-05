package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Conn *mongo.Client
}

func ConnectMongodb() *mongo.Client {

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASS"),
	}
	clientOpts := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT")).SetAuth(credential)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
