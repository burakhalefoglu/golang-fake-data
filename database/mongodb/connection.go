package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Connection struct {
	Conn *mongo.Client
}

var Conn = Connection{}

func ConnectMongodb() *mongo.Client {
	var uri = "mongodb://" + os.Getenv("MONGODB_USER") + ":" + os.Getenv("MONGODB_PASS") + "@" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT") + "/?maxPoolSize=20&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
