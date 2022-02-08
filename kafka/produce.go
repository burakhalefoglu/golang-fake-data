package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("64.227.7.141:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CreateTopic() {
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "64.227.7.141:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func ConvertStructToByteArray(v interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(v)
	return reqBodyBytes.Bytes()
}


func ProduceMesaage(ByteMessage []byte) error{
	writer := newKafkaWriter()
	defer writer.Close()
	
	
	msg := kafka.Message{
		Value: []byte(ByteMessage),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
		
}

