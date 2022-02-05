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
	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS test WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`); err != nil {
		log.Fatal("create keyspace:", err)
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS test.persons (
		id uuid PRIMARY KEY,
		name text,
		gender text,
		address text,
		country text,
		city text,
		continent text,
		region text,
		age bigint,
		married boolean,
		phone text,
		credit_card_number text,
		credit_card_expiration_date_string text,
		credit_card_type text,
		total_spending_gold double,
		total_session_duration double,
		starting_date bigint,
		current_date bigint,
		starting_month bigint,
		current_month bigint,
		total_click_event bigint,
		total_session_count bigint,
		total_score double,
		total_gold bigint)`)
	if err != nil {
		log.Fatal("create table: ", err)
	}

	return &session, err
}
