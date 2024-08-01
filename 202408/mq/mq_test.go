package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	mq := Init(5)
	if mq == nil {
		t.Fatal("Expected MQ instance, got nil")
	}
	if len(mq.topics) != 0 {
		t.Fatalf("Expected empty topics, got %d", len(mq.topics))
	}
	if mq.capacity != 5 {
		t.Fatalf("Expected capacity 5, got %d", mq.capacity)
	}
}

func TestMQClient(t *testing.T) {
	mq := Init(5)
	client := mq.MQClient("topic1", "consumer1")
	if client == nil {
		t.Fatal("Expected Client instance, got nil")
	}
	if client.consumer != "consumer1" {
		t.Fatalf("Expected consumer 'consumer1', got %s", client.consumer)
	}
	if client.topic != "topic1" {
		t.Fatalf("Expected topic 'topic1', got %s", client.topic)
	}
	if cap(client.queue) != 5 {
		t.Fatalf("Expected queue capacity 5, got %d", cap(client.queue))
	}
	if len(mq.topics["topic1"]) != 1 {
		t.Fatalf("Expected 1 client for topic1, got %d", len(mq.topics["topic1"]))
	}
}

func TestSendReceive(t *testing.T) {
	mq := Init(5)
	client := mq.MQClient("topic1", "consumer1")
	msg := Message{Content: "hello"}

	mq.Send("topic1", msg)
	received := client.Receive()
	if received != "hello" {
		t.Fatalf("Expected message 'hello', got %s", received)
	}
}

func TestSendQueueFull(t *testing.T) {
	mq := Init(1)
	client := mq.MQClient("topic1", "consumer1")
	msg1 := Message{Content: "msg1"}
	msg2 := Message{Content: "msg2"}

	mq.Send("topic1", msg1)
	mq.Send("topic1", msg2)

	received := client.Receive()
	if received != "msg1" {
		t.Fatalf("Expected message 'msg1', got %s", received)
	}

	select {
	case msg := <-client.queue:
		t.Fatalf("Expected queue to be empty, but received message: %s", msg.Content)
	default:
		// Queue is empty as expected
	}
}
