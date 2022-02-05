package main

import (
	"github.com/joho/godotenv"
	"golang-fake-data/createFakePerson"
	dataaccess "golang-fake-data/dataaccess/cassandra"
	connection "golang-fake-data/database/cassandra"
	"golang-fake-data/helper"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	defer helper.DeleteHealthFile()
	helper.CreateHealthFile()
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	min := 10
	max := 30
	var dur = rand.Intn(max-min) + min
	time.Sleep(time.Duration(dur) * time.Second)
	session, err1 := connection.ConnectDatabase()
	if err1 != nil {
		log.Fatalln("connection err: ", err1)
		return
	}

	for i := 0; i < 100000000; i++ {
		var fakeData = createFakePerson.CreatePerson()
		log.Println(fakeData)
		var err = dataaccess.InsertData(session, &fakeData)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(fakeData.Name + " added-fake-data")
	}
	log.Println("Finished")
}
