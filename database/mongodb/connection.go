package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Conn *mongo.Client
}

var Conn = Connection{}

func ConnectMongodb()(*mongo.Client){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url = "mongodb://" + os.Getenv("MONGODB_USER") + ":" + os.Getenv("MONGODB_PASS") + "@" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		panic(err)
	}
	return  client
}