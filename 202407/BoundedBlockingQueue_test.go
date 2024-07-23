package main

import (
	"context"
	"testing"
	"time"
)

// 测试Put和Get操作
func TestBoundedBlockingQueue(t *testing.T) {
	queue := NewBoundedBlockingQueue(queueCapacity) // 使用常量创建有界阻塞队列

	// 测试Put操作
	go func() {
		for i := 1; i <= queueCapacity; i++ {
			queue.Put(i)
		}
	}()

	// 确保Put操作正确
	time.Sleep(waitTimeOut) // 等待以确保队列已填满
	if len(queue.items) != queueCapacity {
		t.Errorf("队列长度应为%d，但是是%d", queueCapacity, len(queue.items))
	}

	// 测试Get操作
	go func() {
		for i := 1; i <= queueCapacity; i++ {
			item, _ := queue.Get()
			if item != i {
				t.Errorf("应得元素%d，但得到%v", i, item)
			}
		}
	}()

	// 确保Get操作正确
	time.Sleep(waitTimeOut) // 等待以确保元素已被取出
	if len(queue.items) != 0 {
		t.Errorf("队列应该为空，但是长度为%d", len(queue.items))
	}
}

// 测试生产者和消费者的协同工作
func TestProducerConsumer(t *testing.T) {
	queue := NewBoundedBlockingQueue(queueCapacity) // 使用常量创建有界阻塞队列
	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	// 启动生产者和消费者
	for i := 0; i < producerConsumer; i++ {
		go producer(ctx, queue, i)
	}
	for i := 0; i < producerConsumer; i++ {
		go consumer(ctx, queue, i)
	}

	// 等待context超时
	<-ctx.Done()
}
