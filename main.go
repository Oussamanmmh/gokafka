package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	/*Create a new producer instance
	with the specified configuration
	"bootstrap.servers" specifies the Kafka broker address
	"client.id" specifies the client ID for the producer
	"acks" specifies the acknowledgment level for message delivery
	"all" means the leader will wait for all replicas to acknowledge the message*/
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "myProducer",
		"acks":              "all",
		"security.protocol": "plaintext",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	/*create a channel to handle delivery reports
	"deleveryChan" is a channel of type kafka.Event with a buffer size of 10000
	"topic" is the name of the Kafka topic to which messages will be sent*/

	deleveryChan := make(chan kafka.Event, 10000)
	topic := "HVSE" //name of the kafka topic

	/*produce a message to the specified topic
	"p.Produce" sends the message to the Kafka broker
	*/
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello World !"),
	}, deleveryChan)
	if err != nil {
		log.Fatal(err)
	}
	e := <-deleveryChan

	fmt.Printf("%+v", e)
	fmt.Printf("%+v", p)

}
