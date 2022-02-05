package dataaccess

import (
	"context"
	"golang-fake-data/fakePersonStruct"

	"go.mongodb.org/mongo-driver/mongo"
)

type MDTestDal struct {
	Client *mongo.Client
}

func (m *MDTestDal) Add(data *fakePersonStruct.Person) error {
	ctx := context.Background()
	collection := m.Client.Database("AppneuronTestDatabase").Collection("fakePersons")
	var _, err = collection.InsertOne(ctx, &data)
	if err != nil {
		return err
	}
	return nil
}
