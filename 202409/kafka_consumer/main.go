package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{"192.168.146.128:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalf("Failed to close consumer: %v", err)
		}
	}(consumer)

	consumePartition, err := consumer.ConsumePartition("sep", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer func(consumePartition sarama.PartitionConsumer) {
		err := consumePartition.Close()
		if err != nil {
			log.Fatalf("Failed to close partition consumer: %v", err)
		}
	}(consumePartition)

	for message := range consumePartition.Messages() {
		log.Printf("Received message: %s", string(message.Value))
	}
}
