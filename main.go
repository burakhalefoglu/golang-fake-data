package main

import (
	"github.com/joho/godotenv"
	"golang-fake-data/createFakePerson"
	dataaccess "golang-fake-data/dataaccess/cassandra"
	connection "golang-fake-data/database/cassandra"
	"golang-fake-data/helper"
	"log"
	"runtime"
)

func main() {
	defer helper.DeleteHealthFile()
	helper.CreateHealthFile()
	runtime.MemProfileRate = 0
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	session, err := connection.ConnectDatabase()
	if err != nil {
		log.Fatalln("connection err: ", err)
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
