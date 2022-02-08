package main

import (
	"github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/joho/godotenv"
	"golang-fake-data/createFakePerson"
	dataaccess "golang-fake-data/dataaccess/cassandra"
	connection "golang-fake-data/database/cassandra"
	"golang-fake-data/helper"
	"runtime"
	"golang-fake-data/kafka"
)

func main() {
	defer helper.DeleteHealthFile()
	helper.CreateHealthFile()
	logger.Log.App = "golang-fake-data"

	runtime.MemProfileRate = 0
	if err := godotenv.Load(); err != nil {
		clogger.Error(&logger.Messages{
			"message": "Error loading .env file",
		})
		return
	}
	// session, err := connection.ConnectDatabase()
	// if err != nil {
	// 	clogger.Error(&logger.Messages{
	// 		"connection err: ": err.Error(),
	// 	})
	//}

	

	

	for i := 0; i < 100000000; i++ {
		var fakeData = createFakePerson.CreatePerson()
		var fakeDataByte = kafka.ConvertStructToByteArray(fakeData)
	    if err := kafka.ProduceMesaage(fakeDataByte); err != nil {
			clogger.Error(&logger.Messages{
			"produce error: ": fakeData,
		})}
		clogger.Info(&logger.Messages{
			"produce worked : ": "successed",
		})
		//var err = dataaccess.InsertData(session, &fakeData)
		// if err != nil {
		// 	clogger.Error(&logger.Messages{
		// 		"insert err: ": err.Error(),
		// 	})
		// }
		// clogger.Info(&logger.Messages{
		// 	"person added: ": fakeData,
		// })
	}
	clogger.Info(&logger.Messages{
		"messages: ": "finished all work",
	})
}
