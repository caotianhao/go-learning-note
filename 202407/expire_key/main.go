// 题目：设计一个带过期时间的缓存
// 实现一个有最大容量和过期时间的缓存，实现如下方法：
// Get(key string)(value string, found bool)：获取key，找到就返回true，
// 没找到或者key过期了就返回false，同时实现过期检查
// Put(key string, value sting)：插入或者更新key，超过容量后随机删除一个过期的key
// Delete(key string)：删除一个key
package main

import (
	"fmt"
	"time"
)

const maxCap = 1<<6 - 1

type MyCache struct {
	value    string
	deadTime time.Time
}

type Cache struct {
	doc map[string]MyCache
	cap int
	dur time.Duration
}

// NewCache just new a cache with cap and exp
func NewCache(cap int, dur time.Duration) *Cache {
	return &Cache{
		doc: make(map[string]MyCache),
		cap: cap,
		dur: dur,
	}
}

// Get if cache not found or expire, return "" and false
// and if key exist, return its value
func (c *Cache) Get(key string) (string, bool) {
	res, ok := c.doc[key]
	if !ok || time.Now().After(res.deadTime) {
		if ok {
			c.Delete(key)
		}
		return "", false
	}
	return res.value, true
}

// Put if the len bigger than cap, random delete
func (c *Cache) Put(key string, value string) {
	if len(c.doc) >= c.cap {
		c.randomDelete()
	}
	c.doc[key] = MyCache{
		value:    value,
		deadTime: time.Now().Add(c.dur),
	}
}

// Delete just delete from map
func (c *Cache) Delete(key string) {
	delete(c.doc, key)
}

// randomDelete with small letter and delete a key randomly
func (c *Cache) randomDelete() {
	for key, item := range c.doc {
		if time.Now().After(item.deadTime) {
			c.Delete(key)
			return
		}
	}
	for key := range c.doc {
		c.Delete(key)
		return
	}
}

func main() {
	cache := NewCache(maxCap, 33*time.Second)

	for i := 0; i < maxCap; i++ {
		cache.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i*i*i))
	}

	if v11, ok := cache.Get("k11"); ok {
		fmt.Println("k11:", v11)
	} else {
		fmt.Println("k11 not found or expire")
	}

	cache.Put("k64", "vvvvv")

	if v64, ok := cache.Get("k64"); ok {
		fmt.Println("k64:", v64)
	} else {
		fmt.Println("k64 not found or expire")
	}
}
