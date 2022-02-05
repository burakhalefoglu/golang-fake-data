package dataaccess

import (
	"context"
	"golang-fake-data/fakePersonStruct"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MDTestDal struct {
	Client *mongo.Client
}

func (m *MDTestDal) Add(data *fakePersonStruct.Person) error {

	collection := m.Client.Database("AppneuronTestDatabase").Collection("fakePersons")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var _, err = collection.InsertOne(ctx, &data)

	if err != nil {
		return err
	}
	return nil
}
