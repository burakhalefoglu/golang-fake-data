package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"os"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CreateTopic() {
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", os.Getenv("KAFKA_HOST")+":"+os.Getenv("KAFKA_PORT"))
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
	err := json.NewEncoder(reqBodyBytes).Encode(v)
	if err != nil {
		return nil
	}
	return reqBodyBytes.Bytes()
}

func ProduceMessage(topic string, ByteMessage []byte) error {
	writer := newKafkaWriter(topic)
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
