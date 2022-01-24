package main

import (
	"golang-fake-data/createFakePerson"
	"golang-fake-data/dataaccess"
	"golang-fake-data/database/mongodb"
	"golang-fake-data/helper"
	"log"
	"runtime"
	"time"

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
	var client = mongodb.ConnectMongodb()
	testDal := dataaccess.MDTestDal{
		Client: client,
	}

	for i := 1; i <= 10000000; i++ {
		var fakeData = createFakePerson.CreatePerson()
		testDal.Add(&fakeData)
		time.Sleep(1 * time.Second)
	}
}
