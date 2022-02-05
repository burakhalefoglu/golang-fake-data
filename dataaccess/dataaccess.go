package dataaccess

import (
	"golang-fake-data/database/mongodb"
	"golang-fake-data/fakePersonStruct"
)

type MDTestDal struct {
}

func (m *MDTestDal) Add(data *fakePersonStruct.Person) error {

	client, ctx, cancel, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

	defer mongodb.Close(client, ctx, cancel)
	mongodb.Ping(client, ctx)

	collection := client.Database("AppneuronTestDatabase").Collection("fakePersons")
	var _, errResult = collection.InsertOne(ctx, &data)
	if errResult != nil {
		return errResult
	}
	return nil
}
