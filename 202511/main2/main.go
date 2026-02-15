package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCache struct {
	data map[any]Item
	mu   sync.RWMutex
	sig  chan struct{}
	ttl  time.Duration
	run  bool
}

type Item struct {
	v      any
	expire int64 // nano
}

func NewSafeCache(ttl time.Duration) *SafeCache {
	cache := &SafeCache{
		data: make(map[any]Item),
		mu:   sync.RWMutex{},
		sig:  make(chan struct{}),
		ttl:  ttl,
		run:  true,
	}
	go func() {
		cache.init1()
	}()

	return cache
}

func (c *SafeCache) init1() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.Clean()
		case <-c.sig:
			return
		}
	}
}

func (c *SafeCache) Set(k any, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = Item{
		v:      v,
		expire: time.Now().Add(c.ttl).UnixNano(),
	}
}

func (c *SafeCache) Get(k any) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	res, ok := c.data[k]
	if !ok {
		return nil, false
	}
	if res.expire > time.Now().UnixNano() {
		return nil, false
	}

	return res.v, true
}

func (c *SafeCache) Clean() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()
	for k, v := range c.data {
		if v.expire < now {
			delete(c.data, k)
		}
	}
}

func (c *SafeCache) Close() {
	if c.run {
		close(c.sig)
	}
}

func main() {
	c := NewSafeCache(time.Second)
	defer c.Close()

	c.Set("K1", "V1")
	fmt.Println(c.Get("K1"))

	time.Sleep(2 * time.Second)
	fmt.Println(c.Get("K1"))
}
