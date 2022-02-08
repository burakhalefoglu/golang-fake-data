package kafka

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"kafka/personStruct"

// 	"github.com/segmentio/kafka-go"
// )

// func ReadKafka(obj *personStruct.Person) {
// 	conf := kafka.ReaderConfig{
// 		Brokers:  []string{"64.227.7.141:9092"},
// 		Topic:    "my-topic",
// 		GroupID:  "g1",
// 		MaxBytes: 1000,
// 	}

// 	reader := kafka.NewReader(conf)

// 	for {
// 		m, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Println("error occured", err)
// 			continue
// 		}
// 		fmt.Println("message is : ", string(m.Value))
// 		if err := json.Unmarshal(m.Value, obj); err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(obj)
// 	}

// }
