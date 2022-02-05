package main

import (
	"golang-fake-data/createFakePerson"
	dataaccess "golang-fake-data/dataaccess/cassandra"
	"golang-fake-data/helper"
	"log"
	"runtime"

	"github.com/joho/godotenv"
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

	for i := 0; i < 100000000; i++ {
		var fakeData = createFakePerson.CreatePerson()
		log.Println(fakeData)
		var err = dataaccess.InsertData(&fakeData)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(fakeData.Name + " Added")
	}
	log.Println("Finished")
}
