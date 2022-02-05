package dataaccess

import (
	connection "golang-fake-data/database/cassandra"
	"golang-fake-data/fakePersonStruct"
	"log"
)

func InsertData(data *fakePersonStruct.Person) error {
	session, err := connection.ConnectDatabase()
	q := session.Query(fakePersonStruct.PersonTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer session.Close()
	return nil

}
