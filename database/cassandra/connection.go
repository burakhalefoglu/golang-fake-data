package connection

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
)

func ConnectDatabase() (*gocqlx.Session, error) {

	cluster := gocql.NewCluster("k8ssandra-cassandra.k8ssandra.svc.cluster.local")
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "test*12",
	}
	cluster.Consistency = gocql.Quorum
	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &session, err
}
