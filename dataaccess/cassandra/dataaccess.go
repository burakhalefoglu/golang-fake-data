package dataaccess

import (
	"fmt"
	connection "golang-fake-data/database/cassandra"
	"golang-fake-data/fakePersonStruct"
	"log"
)

func InsertData(data *fakePersonStruct.Person) error {
	session, err := connection.ConnectDatabase()
	if err != nil {
		log.Fatalln("create keyspace err: ", err)
		return err
	}

	err1 := session.ExecStmt(fmt.Sprintf(`CREATE KEYSPACE  IF NOT EXISTS  %s WITH replication = {'class' : 'SimpleStrategy', 'replication_factor' : %d}`, "AppneuronTestDatabase", 3))
	if err1 != nil {
		log.Fatalln("create keyspace err: ", err1)
		return err1
	}

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
