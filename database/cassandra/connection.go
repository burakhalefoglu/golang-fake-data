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
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Create table

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS Person (
		id uuid PRIMARY KEY,
		"name text",
		"gender text",
		"address text",
		"country text",
		"city text",
		"continent text",
		"region text",
		"age bigint",
		"married boolean",
		"phone text",
		"credit_card_number text",
		"credit_card_expiration_date_string text",
		"credit_card_type text",
		"total_spending_gold double",
		"total_session_duration double",
		"starting_date bigint",
		"current_date bigint",
		"starting_month bigint",
		"current_month bigint",
		"total_click_event bigint",
		"total_session_count bigint",
		"total_score double",
		"total_gold bigint",)`)
	if err != nil {
		log.Fatal("create table: ", err)
	}

	return &session, err
}
