package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 导入 Redis Go 客户端库
	"github.com/google/uuid"       // 导入 UUID 库，用于生成唯一的锁标识符
)

var ctx1 = context.Background() // 创建一个背景上下文，用于 Redis 操作

var redisClient *redis.Client // Redis 客户端实例

// initialize 函数用于初始化 Redis 客户端
// 参数 redisAddr: Redis 服务器的地址和端口，例如 "localhost:6379"
func initialize(redisAddr string) {
	// 创建一个 Redis 客户端并配置连接选项
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr, // 设置 Redis 服务器的地址和端口
	})
}

// acquireLock 函数用于尝试获取分布式锁
// 参数 lockName: 锁的名称
// 参数 expire: 锁的过期时间
// 返回值: 锁的唯一标识符和可能出现的错误
func acquireLock(lockName string, expire time.Duration) (string, error) {
	// 生成一个唯一的锁标识符
	lockValue := uuid.New().String()

	// 尝试设置锁的值，使用 SETNX 命令仅在锁不存在时设置
	// 设置锁的过期时间
	success, err := redisClient.SetNX(ctx1, lockName, lockValue, expire).Result()
	if err != nil {
		return "", err // 如果设置失败，返回错误
	}
	if !success {
		return "", nil // 如果锁已经存在，返回空字符串，表示无法获取锁
	}
	return lockValue, nil // 返回锁的唯一标识符
}

// releaseLock 函数用于释放分布式锁
// 参数 lockName: 锁的名称
// 参数 lockValue: 锁的唯一标识符
// 返回值: 布尔值表示是否成功释放锁和可能出现的错误
func releaseLock(lockName, lockValue string) (bool, error) {
	// Lua 脚本用于确保只有持有锁的客户端才能释放锁
	// 脚本逻辑: 如果当前锁值与提供的值匹配，则删除锁
	script := `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end
    `

	// 执行 Lua 脚本
	result, err := redisClient.Eval(ctx1, script, []string{lockName}, lockValue).Result()
	if err != nil {
		return false, err // 如果执行脚本失败，返回错误
	}

	// 返回是否成功释放锁
	return result.(int64) == 1, nil
}

// main 函数是程序的入口点
func main() {
	// 初始化 Redis 客户端，连接到指定的 Redis 服务器
	initialize("localhost:6379")

	// 尝试获取锁
	lockValue, err := acquireLock("my_lock", 10*time.Second)
	if err != nil {
		fmt.Println("Failed to acquire lock:", err) // 打印错误信息
		return
	}
	if lockValue == "" {
		fmt.Println("Lock is already held by another process") // 锁已经被其他进程持有
		return
	}
	fmt.Println("Lock acquired successfully") // 成功获取锁

	// 模拟锁的使用
	time.Sleep(5 * time.Second)

	// 尝试释放锁
	success, err := releaseLock("my_lock", lockValue)
	if err != nil {
		fmt.Println("Failed to release lock:", err) // 打印错误信息
		return
	}
	if success {
		fmt.Println("Lock released successfully") // 成功释放锁
	} else {
		fmt.Println("Failed to release lock, it might have expired or been held by another process") // 释放锁失败
	}
}
