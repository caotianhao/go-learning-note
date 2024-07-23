package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 常量定义
const (
	parkingLotCapacity = 10               // 停车场容量
	secondlyRate       = 2                // 每秒停车费用
	maxStay            = 10 * time.Second // 每辆车最大停车时间，10小时换算成程序时间
)

// ParkingLot 表示停车场结构体
type ParkingLot struct {
	capacity    int           // 停车场最大容量
	occupied    int           // 当前已占用车位数
	mu          sync.Mutex    // 互斥锁，用于保护对停车场状态的访问
	sem         chan struct{} // 通道用于模拟信号量
	totalIncome float64       // 停车场总收入
}

// NewParkingLot 创建一个新的停车场实例
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		sem:      make(chan struct{}, capacity), // 创建带缓冲的通道，容量为停车场大小
	}
}

// EnterCar 尝试将汽车进入停车场
func (pl *ParkingLot) EnterCar(carID int) {
	pl.sem <- struct{}{} // 请求停车位，如果没有空位则会阻塞
	pl.mu.Lock()
	defer pl.mu.Unlock()
	pl.occupied++
	fmt.Printf("汽车 %d 进入停车场。当前占用车位: %d/%d\n", carID, pl.occupied, pl.capacity)
}

// ExitCar 处理汽车离开停车场
func (pl *ParkingLot) ExitCar(carID int, parkingDuration time.Duration) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	pl.occupied--
	fmt.Printf("汽车 %d 离开停车场。停车时间: %v，当前占用车位: %d/%d\n", carID, parkingDuration, pl.occupied, pl.capacity)
	tmpIncome := parkingDuration.Seconds() / secondlyRate
	//fmt.Println("aaa", tmpIncome)
	pl.totalIncome += tmpIncome // 计算停车费用
	<-pl.sem                    // 释放停车位
}

// GetTotalIncome 获取停车场总收入
func (pl *ParkingLot) GetTotalIncome() float64 {
	return pl.totalIncome
}

// RandomCarGenerator 每秒钟随机产生0到5辆汽车
func RandomCarGenerator(pl *ParkingLot, stopChan chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	carID := 0
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			numCars := rand.Intn(6) // 生成0到5之间的随机数
			for i := 0; i < numCars; i++ {
				carID++
				go func(id int) {
					// 模拟汽车到达停车场并停放
					pl.EnterCar(id)
					// 模拟汽车停放时间
					startTime := time.Now()
					time.Sleep(time.Duration(rand.Intn(int(maxStay))))
					pl.ExitCar(id, time.Since(startTime))
				}(carID)
			}
		}
	}
}
