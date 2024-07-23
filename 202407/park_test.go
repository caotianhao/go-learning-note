package main

import (
	"testing"
	"time"
)

// TestParkingLot 测试停车场的功能
func TestParkingLot(t *testing.T) {
	pl := NewParkingLot(parkingLotCapacity) // 创建停车场实例
	stopChan := make(chan struct{})         // 创建停止信号通道
	defer close(stopChan)                   // 确保测试结束时关闭通道

	go RandomCarGenerator(pl, stopChan) // 启动随机汽车生成器

	// 等待一定时间以便汽车完成停车
	time.Sleep(11 * time.Second) // 确保所有任务完成

	totalIncome := pl.GetTotalIncome() // 获取总收入

	// 验证总收入是否符合预期
	if totalIncome < 0 {
		t.Errorf("期望的总收入为大于等于0，但实际为 %.2f", totalIncome)
	} else {
		t.Logf("停车场总收入: %.2f 元", totalIncome)
	}
}
