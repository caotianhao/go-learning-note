package workpool

import (
	"sync"
	"time"
)

const (
	workPoolSize   = 5                      // 工作协程数量
	taskQueueSize  = 10 * workPoolSize      // 任务队列大小
	numberOfTask   = 2 * taskQueueSize      // 任务总数
	taskSleepTime  = 375 * time.Millisecond // 模拟任务执行时间
	poolCloseDelay = 5 * time.Second        // 关闭工作池前的延迟时间
)

// WorkPool 工作池结构体
type WorkPool struct {
	workPoolSize int // 工作协程数量
	taskQueue    chan func()
	wg           sync.WaitGroup
	closeChan    chan struct{}
}

// NewWorkPool 创建一个新工作池
func NewWorkPool(workPoolSize int) *WorkPool {
	return &WorkPool{
		workPoolSize: workPoolSize,
		taskQueue:    make(chan func(), taskQueueSize),
		closeChan:    make(chan struct{}),
	}
}

// Start 启动工作池，启动指定数量的工作协程
func (wp *WorkPool) Start() {
	wp.wg.Add(wp.workPoolSize)
	for i := 0; i < wp.workPoolSize; i++ {
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case task := <-wp.taskQueue:
					task()
				case <-wp.closeChan:
					return
				}
			}
		}()
	}
}

// SubmitTask 提交任务
func (wp *WorkPool) SubmitTask(task func()) {
	wp.taskQueue <- task
}

// Close 优雅关闭工作池
func (wp *WorkPool) Close() {
	close(wp.closeChan)
	wp.wg.Wait()
}
