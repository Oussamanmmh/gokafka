package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "HVSE" //name of the kafka topic

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "myDataGroup",
		"enable.auto.commit": "true",
		"security.protocol":  "plaintext",
	})
	/*Error handling is important*/
	if err != nil {
		log.Fatal("fail to create consumer .", err)
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal("fail to create subscribe .", err)
	}
	for {
		ev := consumer.Poll(100) // 100ms timeout for polling events
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("data team reading order : %s\n", string(e.Value))
		case kafka.Error:
			fmt.Printf("Error: %v\n", e)

		}
	}

}
