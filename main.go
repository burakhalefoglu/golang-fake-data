package main

import (
	"golang-fake-data/createFakePerson"
	"golang-fake-data/dataaccess"
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

	testDal := dataaccess.MDTestDal{}

	for i := 0; i < 100000000; i++ {
		var fakeData = createFakePerson.CreatePerson()
		log.Println(fakeData)
		err := testDal.Add(&fakeData)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(fakeData.Name + " Added")
	}
	log.Println("Finished")
}
