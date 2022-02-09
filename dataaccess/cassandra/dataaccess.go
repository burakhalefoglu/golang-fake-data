package dataaccess

import (
	"golang-fake-data/fakePersonStruct"

	"github.com/scylladb/gocqlx/v2"
)

func InsertData(session *gocqlx.Session, data *fakePersonStruct.Person) error {
	q := session.Query(fakePersonStruct.PersonTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}
