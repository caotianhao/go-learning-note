// go 实现 ping 操作

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"time"
)

var (
	timeout      int64
	size         int
	count        int
	typ          uint8 = 8
	code         uint8 = 0
	sendCount    int
	successCount int
	failCount    int
	minTime      int64 = math.MaxInt32
	maxTime      int64
	totalTime    int64
)

// ICMP 报文，必须严格按照报文结构定义
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "等待每次回复的超时时间(毫秒)")
	flag.IntVar(&size, "l", 32, "发送缓冲区大小")
	flag.IntVar(&count, "n", 4, "要发送的回显请求数")
	flag.Parse()
}

func checkSumTrans(data []byte) uint16 {
	l := len(data)
	index := 0
	var sum uint32 = 0
	for l > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		l -= 2
		index += 2
	}
	if l != 0 {
		sum += uint32(data[index])
	}
	high16 := sum >> 16
	for high16 != 0 {
		sum = high16 + uint32(uint16(sum))
		high16 = sum >> 16
	}
	return uint16(^sum)
}

func main() {
	// 拿到命令行参数
	getCommandArgs()

	// ip 地址为命令行参数的最后一个
	dstIP := os.Args[len(os.Args)-1]

	// 建立链接，使用 ICMP 协议
	conn, err := net.DialTimeout("ip:icmp", dstIP, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		fmt.Println("net.DialTimeout failed", err)
		return
	}

	// defer
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close() failed", err)
			return
		}
	}(conn)

	// 输出第一行
	fmt.Printf("正在 Ping %s [%s] 具有 %d 字节的数据\n", dstIP, conn.RemoteAddr(), size)

	// 按照所选的 count 数量进行 count 次循环
	for i := 0; i < count; i++ {
		// 发送量++
		sendCount++
		// icmp 参数
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			CheckSum:    0,
			ID:          0,
			SequenceNum: 1,
		}

		data := make([]byte, size)
		var buffer bytes.Buffer
		err = binary.Write(&buffer, binary.BigEndian, icmp)
		if err != nil {
			fmt.Println("binary.Write fail", err)
			return
		}

		buffer.Write(data)
		data = buffer.Bytes()

		checkSum := checkSumTrans(data)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)

		timeBegin := time.Now()
		err = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		if err != nil {
			return
		}

		_, err := conn.Write(data)
		if err != nil {
			failCount++
			continue
		}

		buf := make([]byte, 65535)
		r, err := conn.Read(buf)
		if err != nil {
			failCount++
			continue
		}

		successCount++

		ts := time.Since(timeBegin).Milliseconds()
		if minTime > ts {
			minTime = ts
		}
		if maxTime < ts {
			maxTime = ts
		}
		totalTime += ts

		fmt.Printf("来自 %d.%d.%d.%d 的回复：字节=%d 时间=%d ms TTL=%d\n",
			buf[12], buf[13], buf[14], buf[15], r-28, ts, buf[8])
		time.Sleep(time.Second)
	}
	
	fmt.Printf("%s 的 Ping 统计信息: 数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%.2f%%丢失)，\n"+
		"往返行程的估计时间(以毫秒为单位): 最短 = %d ms，最长 = %d ms，平均 = %d ms\n",
		conn.RemoteAddr(), sendCount, successCount, failCount,
		float64(failCount)/float64(sendCount)*100, minTime, maxTime, totalTime/int64(sendCount))
}
