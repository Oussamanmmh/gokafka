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
		"group.id":           "myGroup",
		"enable.auto.commit": true,
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
			fmt.Printf("proccessing order: %s\n", string(e.Value))
			_, err := consumer.CommitMessage(e)
			if err != nil {
				log.Fatal("fail to commit message .", err)
			}
		case kafka.Error:
			fmt.Printf("Error: %v\n", e)

		}
	}

}
