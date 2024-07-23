package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	producerConsumer = 12 // 生产者和消费者的数量
	queueCapacity    = 16 // 队列的容量
)

const (
	waitTimeOut = 500 * time.Millisecond // 超时时间
	ctxTime     = 10 * time.Second       // ctx控制时间
)

// BoundedBlockingQueue 定义了一个有界阻塞队列
type BoundedBlockingQueue struct {
	items    []interface{} // 存储队列中的元素
	capacity int           // 队列的容量
	mu       sync.Mutex    // 互斥锁，保护共享资源
	notEmpty *sync.Cond    // 用于等待队列非空的条件变量
	notFull  *sync.Cond    // 用于等待队列非满的条件变量
}

// NewBoundedBlockingQueue 创建一个新的有界阻塞队列
func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {
	if capacity <= 0 { // 如果容量小于等于0，抛出异常
		panic("容量必须大于0")
	}
	q := &BoundedBlockingQueue{
		items:    make([]interface{}, 0, capacity), // 初始化队列，容量为指定值
		capacity: capacity,                         // 设置队列容量
	}
	q.notEmpty = sync.NewCond(&q.mu) // 初始化notEmpty条件变量
	q.notFull = sync.NewCond(&q.mu)  // 初始化notFull条件变量
	return q
}

// Put 将一个元素放入队列，如果队列已满则阻塞
func (q *BoundedBlockingQueue) Put(item interface{}) {
	q.mu.Lock()                      // 加锁，保护共享资源
	defer q.mu.Unlock()              // 解锁，保证在函数退出时解锁
	for len(q.items) == q.capacity { // 如果队列已满
		fmt.Println("队列已满，Put操作阻塞中...") // 打印提示信息
		q.notFull.Wait()                // 等待队列有空位
	}
	q.items = append(q.items, item) // 将新元素加入队列
	fmt.Printf("元素%v已加入队列\n", item) // 打印提示信息
	q.notEmpty.Signal()             // 通知消费者队列不为空
}

// Get 从队列中取出一个元素，如果队列为空则阻塞
func (q *BoundedBlockingQueue) Get() (item interface{}, err error) {
	q.mu.Lock()             // 加锁，保护共享资源
	defer q.mu.Unlock()     // 解锁，保证在函数退出时解锁
	for len(q.items) == 0 { // 如果队列为空
		fmt.Println("队列为空，Get操作阻塞中...") // 打印提示信息
		q.notEmpty.Wait()               // 等待队列有元素
	}
	item = q.items[0]                // 从队列头部取出元素
	q.items = q.items[1:]            // 移除取出的元素
	fmt.Printf("元素%v已从队列取出\n", item) // 打印提示信息
	q.notFull.Signal()               // 通知生产者队列不满
	return item, nil
}

// producer 生产者函数，持续向队列中放入元素
func producer(ctx context.Context, queue *BoundedBlockingQueue, id int) {
	for {
		select {
		case <-ctx.Done(): // 如果context被取消，停止生产
			fmt.Printf("生产者%d停止\n", id) // 打印提示信息
			return
		default:
			queue.Put(id)                      // 将生产的元素放入队列
			fmt.Printf("生产者%d生产了%d\n", id, id) // 打印生产的元素
			time.Sleep(waitTimeOut)            // 模拟生产时间
		}
	}
}

// consumer 消费者函数，持续从队列中取出元素
func consumer(ctx context.Context, queue *BoundedBlockingQueue, id int) {
	for {
		select {
		case <-ctx.Done(): // 如果context被取消，停止消费
			fmt.Printf("消费者%d停止\n", id) // 打印提示信息
			return
		default:
			item, err := queue.Get() // 从队列中取出元素
			if err != nil {
				fmt.Printf("消费者%d错误: %v\n", id, err) // 打印错误信息
				return
			}
			fmt.Printf("消费者%d消费了%v\n", id, item) // 打印取出的元素
			time.Sleep(waitTimeOut)              // 模拟消费时间
		}
	}
}
