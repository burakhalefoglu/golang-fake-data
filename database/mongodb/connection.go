package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MONGODB'YE BAĞLANIYORUZ
func Connect() (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASS"),
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	clientOpts := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT")).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	return client, ctx, cancel, err
}

//BAĞLANTI TESTİ YAPIYORUZ
func Ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Println("Bağlanti basarili!")
	return nil
}

//MONGODB BAĞLANTISINI SONLANDIRIYORUZ
func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
