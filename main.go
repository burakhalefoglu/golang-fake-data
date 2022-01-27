package main

import (
	"golang-fake-data/createFakePerson"
	"golang-fake-data/dataaccess"
	"golang-fake-data/database/mongodb"
	"golang-fake-data/helper"
	"log"
	"runtime"
	"github.com/joho/godotenv"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go func(i int) {
			defer wg.Done()
				var fakeData = createFakePerson.CreatePerson()
				err := testDal.Add(&fakeData)
				if err != nil {
					log.Fatal(err)
				}
		}(i)
	}
	wg.Wait()
}
