package connection

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
	"os"
)

func ConnectDatabase() (*gocqlx.Session, error) {

	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	//cluster.ConnectTimeout = time.Second * 30
	//cluster.Timeout = time.Second * 30
	//cluster.NumConns = 1000
	//cluster.ReconnectInterval = time.Second * 30
	//cluster.SocketKeepalive = 0
	//cluster.DisableInitialHostLookup = true
	//cluster.IgnorePeerAddr = true
	//cluster.Events.DisableNodeStatusEvents = true
	//cluster.Events.DisableTopologyEvents = true
	//cluster.Events.DisableSchemaEvents = true
	//cluster.MaxRoutingKeyInfo = 50000
	//cluster.PageSize = 50000
	//cluster.WriteCoalesceWaitTime = time.Second * 30
	//cluster.ReconnectionPolicy = &gocql.ConstantReconnectionPolicy{MaxRetries: 5000, Interval: 5 * time.Second}
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASS"),
	}
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS AppneuronTestDatabase WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`); err != nil {
		log.Fatal("create keyspace:", err)
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS AppneuronTestDatabase.fakePersons (
		id text PRIMARY KEY,
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
		total_gold bigint,
		total_skill_count bigint,
		swipe_right_count bigint,
		swipe_left_count bigint,
		swipe_down_count bigint,
		swipe_up_count bigint,
		start_cor double,
		finish_cor double,
		remain_life double,
		adv_type text,
		product_type text,
		device_type bigint,
		graphics_device_type bigint,
		graphics_memory_size bigint,
		operating_system bigint,
		processor_count bigint,
		processor_type bigint,
		system_memory_size bigint,
		first_session_year_of_day bigint,
		first_session_year bigint,
		first_session_week_day bigint,
		first_session_hour bigint,
		first_session_duration bigint,
		first_session_minute bigint,
		second_session_hour bigint,
		second_session_duration bigint,
		second_session_minute bigint,
		third_session_hour bigint,
		third_session_duration bigint,
		third_session_minute bigint,
		fourth_session_hour bigint,
		fourth_session_duration bigint,
		fourth_session_minute bigint,
		fifth_session_hour bigint,
		fifth_session_duration bigint,
		fifth_session_minute bigint,
		penultimate_session_hour bigint,
		penultimate_session_duration bigint,
		penultimate_session_minute bigint,
		last_session_year_of_day bigint,
		last_session_year bigint,
		last_session_hour bigint,
		last_session_duration bigint,
		last_session_minute bigint,
		first_half_hour_total_session_count bigint,
		first_half_hour_total_session_duration bigint,
		first_hour_total_session_count bigint,
		first_hour_total_session_duration bigint,
		first_two_hour_total_session_count bigint,
		first_two_hour_total_session_duration bigint,
		first_three_hour_total_session_count bigint,
		first_three_hour_total_session_duration bigint,
		first_six_hour_total_session_count bigint,
		first_six_hour_total_session_duration bigint,
		first_twelve_hour_total_session_count bigint,
		first_twelve_hour_total_session_duration bigint,
		total_session_day bigint,
		total_session_hour bigint,
		total_session_minute bigint,
		first_day_total_session_count bigint,
		first_day_total_session_duration bigint,
		second_day_total_session_count bigint,
		second_day_total_session_duration bigint,
		third_day_total_session_count bigint,
		third_day_total_session_duration bigint,
		fourth_day_total_session_count bigint,
		fourth_day_total_session_duration bigint,
		fifth_day_total_session_count bigint,
		fifth_day_total_session_duration bigint,
		sixth_day_total_session_count bigint,
		sixth_day_total_session_duration bigint,
		seventh_day_total_session_count bigint,
		seventh_day_total_session_duration bigint,
		min_session_duration bigint,
		max_session_duration bigint,
		daily_average_session_count double,
		daily_average_session_duration double,
		session_based_average_session_duration double,
		daily_average_session_count_minus_first_day_session_count double,
		daily_average_session_duration_minus_first_day_session_duration double,
		session_based_average_session_duration_minus_first_session_duration double,
		session_based_average_session_duration_minus_last_session_duration double,
		sunday_session_count bigint,
		monday_session_count bigint,
		tuesday_session_count bigint,
		wednesday_session_count bigint,
		thursday_session_count bigint,
		friday_session_count bigint,
		saturday_session_count bigint,
		am_session_count bigint,
		pm_session_count bigint,
		session_zero_to_five_hour_count bigint,
		session_six_to_eleven_hour_count bigint,
		session_twelve_to_seventeen_hour_count bigint,
		session_eighteen_to_twenty_three_hour_count bigint)`)
	if err != nil {
		log.Fatal("create table: ", err)
	}

	return &session, err
}
