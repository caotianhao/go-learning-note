package main

// 基于 channel 实现一个简易的队列中间件
// 初始化：支持设置队列的容量
// 实现多 topic 管理，每个 topic 支持多个消费者（简化操作，生产者只有一个）
// 分别实现生产者和消费者逻辑
// 当队列被占满时，生产者不阻塞，丢弃消息

import (
	"fmt"
)

// Message 结构体定义，表示队列中传递的消息
type Message struct {
	Content string // 消息内容
}

// MsgQueue 是一个消息队列类型，基于通道实现
type MsgQueue chan Message

// Client 结构体定义，表示消费者
type Client struct {
	consumer string   // 消费者名称
	topic    string   // 主题名称
	queue    MsgQueue // 消息队列
}

// MQ 结构体定义，表示消息队列管理器
type MQ struct {
	topics   map[string][]*Client // 主题到客户端列表的映射
	capacity int                  // 队列容量
}

// Init 初始化消息队列管理器，并指定队列的容量
func Init(capacity int) *MQ {
	return &MQ{
		topics:   make(map[string][]*Client), // 初始化主题到客户端列表的映射
		capacity: capacity,                   // 设置队列容量
	}
}

// MQClient 创建一个新的消费者，并将其加入到指定主题的客户端列表中
func (m *MQ) MQClient(topic, consumer string) *Client {
	queue := make(MsgQueue, m.capacity) // 创建指定容量的消息队列
	client := &Client{
		consumer: consumer,
		topic:    topic,
		queue:    queue,
	}
	m.topics[topic] = append(m.topics[topic], client) // 将客户端加入到主题的客户端列表中
	return client
}

// Send 向指定主题发送消息，会被所有订阅了该主题的客户端接收
func (m *MQ) Send(topic string, msg Message) {
	clients, ok := m.topics[topic] // 获取订阅了该主题的所有客户端
	if !ok {
		fmt.Printf("未找到主题 '%s'\n", topic) // 主题不存在的情况下输出错误信息
		return
	}

	for _, client := range clients {
		select {
		case client.queue <- msg: // 尝试向客户端的队列发送消息
			fmt.Printf("消息已发送到主题 '%s' 中的消费者 '%s'\n", topic, client.consumer)
		default:
			fmt.Printf("队列已满，消息丢弃，主题 '%s' 中的消费者 '%s'\n", topic, client.consumer)
		}
	}
}

// Receive 从客户端队列中接收消息
func (c *Client) Receive() string {
	msg := <-c.queue // 从客户端的队列中接收消息
	return msg.Content
}

//func main() {
//	mq := Init(10) // 初始化队列容量为 10 的消息队列管理器
//
//	client1 := mq.MQClient("topic1", "consumer1") // 创建主题1的消费者1
//	client2 := mq.MQClient("topic2", "consumer2") // 创建主题2的消费者2
//
//	mq.Send("topic1", Message{Content: "你好，主题1的消息！"}) // 发送消息到主题1
//	mq.Send("topic2", Message{Content: "你好，主题2的消息！"}) // 发送消息到主题2
//
//	fmt.Println("消费者1收到消息:", client1.Receive()) // 消费者1接收消息
//	fmt.Println("消费者2收到消息:", client2.Receive()) // 消费者2接收消息
//}
