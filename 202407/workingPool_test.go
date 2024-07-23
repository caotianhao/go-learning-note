package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// TestWorkerPool 测试 WorkerPool 的功能
func TestWorkerPool(t *testing.T) {
	// 创建一个新的工作池，工作协程数量由 workerPoolSize 常量指定
	wp := NewWorkerPool(workerPoolSize)
	// 启动工作池，这将启动所有工作协程
	wp.Start()

	// 使用 int32 变量来记录已完成的任务数量，atomic 包提供了原子操作以保证线程安全
	var completedTasks int32

	// 提交多个任务到工作池
	for i := 1; i <= numberOfTask; i++ {
		// 将循环变量 i 传递到闭包中需要创建一个局部副本
		j := i
		wp.SubmitTask(func() {
			// 打印任务的执行信息
			fmt.Printf("执行任务%d\n", j)
			// 原子操作，安全地增加任务计数，防止并发操作造成的问题
			atomic.AddInt32(&completedTasks, 1)
			// 模拟任务执行时间，实际应用中任务的执行时间可能会有所不同
			time.Sleep(taskSleepTime)
		})
	}

	// 等待所有任务有足够的时间执行完成，确保工作池在关闭前完成所有任务
	time.Sleep(poolCloseDelay)
	// 关闭工作池，等待所有工作协程完成任务
	wp.Close()

	// 检查已完成的任务数量是否与预期的数量相符
	if completedTasks != numberOfTask {
		// 如果实际完成的任务数量与预期不符，则报告测试失败
		t.Errorf("期望完成的任务数量为%d，但实际为%d", numberOfTask, completedTasks)
	} else {
		// 如果任务数量匹配，则测试通过，输出成功信息
		fmt.Println("所有任务已完成，工作池已关闭！")
	}
}
