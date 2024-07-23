package main

import (
	"sync"
	"time"
)

const (
	workerPoolSize = 5                      // 工作协程数量
	taskQueueSize  = 10 * workerPoolSize    // 任务队列大小
	numberOfTask   = 2 * taskQueueSize      // 任务总数
	taskSleepTime  = 200 * time.Millisecond // 模拟任务执行时间
	poolCloseDelay = 3 * time.Second        // 关闭工作池前的延迟时间
)

// WorkerPool 表示工作池结构体
type WorkerPool struct {
	workerPoolSize int            // 工作协程数量
	taskQueue      chan func()    // 任务队列
	wg             sync.WaitGroup // 等待组，用于等待所有工作协程完成
	closeChan      chan struct{}  // 用于通知工作协程退出的信号通道
}

// NewWorkerPool 创建一个新的工作池
func NewWorkerPool(workerPoolSize int) *WorkerPool {
	return &WorkerPool{
		workerPoolSize: workerPoolSize,
		taskQueue:      make(chan func(), taskQueueSize), // 假设队列大小是工作协程数量的10倍
		closeChan:      make(chan struct{}),
	}
}

// Start 启动工作池，启动指定数量的工作协程
func (wp *WorkerPool) Start() {
	wp.wg.Add(wp.workerPoolSize) // 设置等待组计数
	for i := 0; i < wp.workerPoolSize; i++ {
		go func() {
			defer wp.wg.Done() // 当协程退出时减少等待组计数
			for {
				select {
				case task := <-wp.taskQueue: // 从任务队列中取出任务
					task() // 执行任务
				case <-wp.closeChan: // 收到关闭信号时退出
					return
				}
			}
		}()
	}
}

// SubmitTask 提交任务到工作池
func (wp *WorkerPool) SubmitTask(task func()) {
	wp.taskQueue <- task // 将任务放入任务队列
}

// Close 优雅关闭工作池，等待所有任务完成
func (wp *WorkerPool) Close() {
	close(wp.closeChan) // 关闭信号通道，通知所有工作协程退出
	wp.wg.Wait()        // 等待所有工作协程完成
}
