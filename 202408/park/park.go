package park

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	parkCap     = 10
	parkRate    = 2
	maxStayTime = 10 * time.Second
)

type Park struct {
	mu         sync.Mutex
	ch         chan struct{}
	cap        int
	used       int
	totalMoney float64
}

func NewPark(cap int) *Park {
	return &Park{
		cap: cap,
		ch:  make(chan struct{}, cap),
	}
}

// EnterCar 去停车
func (p *Park) EnterCar(carID int) {
	p.ch <- struct{}{} // 没有空位就阻塞
	p.mu.Lock()
	defer p.mu.Unlock()
	p.used++
	fmt.Printf("汽车 %d 来了，当前占用车位: %d/%d\n", carID, p.used, p.cap)
}

// ExitCar 处理汽车离开停车场
func (p *Park) ExitCar(carID int, tt time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.used--
	fmt.Printf("汽车 %d 走了，停了 %v，当前占用车位: %d/%d\n", carID, tt, p.used, p.cap)
	tmpMoney := tt.Seconds() / parkRate
	//fmt.Println("aaa", tmpMoney)
	p.totalMoney += tmpMoney // 计算停车费用
	<-p.ch                   // 释放停车位
}

// GetTotalMoney 赚钱啦
func (p *Park) GetTotalMoney() float64 {
	return p.totalMoney
}

// RandomPark 随机停车
func RandomPark(p *Park, stopCh chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	carID := 0
	for {
		select {
		case <-stopCh:
			//fmt.Println("stopCh return")
			return
		case <-ticker.C:
			numCars := rand.Intn(6)
			for i := 0; i < numCars; i++ {
				//fmt.Println("bbbbb", carID)
				carID++
				go func(id int) {
					p.EnterCar(id)
					startTime := time.Now()
					time.Sleep(time.Duration(rand.Intn(int(maxStayTime))))
					p.ExitCar(id, time.Since(startTime))
				}(carID)
			}
		}
		//fmt.Println("Done!!!!!!!!!!!!!!!!")
	}
}
