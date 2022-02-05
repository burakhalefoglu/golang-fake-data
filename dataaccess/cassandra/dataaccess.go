package dataaccess

import (
	"github.com/scylladb/gocqlx/v2"
	"golang-fake-data/fakePersonStruct"
	"log"
)

func InsertData(session *gocqlx.Session, data *fakePersonStruct.Person) error {
	q := session.Query(fakePersonStruct.PersonTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	//defer session.Close()
	return nil
}
