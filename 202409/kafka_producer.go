package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"192.168.146.128:9092"}, config)
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
				Topic: "sep",
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
