package workpool

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkPool(t *testing.T) {
	wp := NewWorkPool(workPoolSize)
	wp.Start()

	var completedTasks int32

	for i := 1; i <= numberOfTask; i++ {
		j := i
		wp.SubmitTask(func() {
			fmt.Printf("执行任务 %d\n", j)
			atomic.AddInt32(&completedTasks, 1)
			// 模拟任务执行时间
			time.Sleep(taskSleepTime)
		})
	}

	time.Sleep(poolCloseDelay)
	wp.Close()

	if completedTasks != numberOfTask {
		t.Errorf("期望完成的任务数量为 %d，但实际为 %d", numberOfTask, completedTasks)
	} else {
		// 如果任务数量匹配，则测试通过，输出成功信息
		fmt.Println("所有任务已完成，工作池已关闭！")
	}
}
