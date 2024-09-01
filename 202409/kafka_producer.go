package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"

	"gogogo/202409/kafka"
)

func main() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(kafka.IpLists, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Fatalf("Failed to close producer: %v", err)
		}
	}(producer)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")

		if scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				break
			}

			message := &sarama.ProducerMessage{
				Topic: kafka.Topic,
				Value: sarama.StringEncoder(line),
			}

			_, _, err := producer.SendMessage(message)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
