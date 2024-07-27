package main

import (
	"fmt"
	"sync"
)


type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// SafeCounterのメソッドであるIncは、カウンタを安全にインクリメントします。
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// LockがUnlockされるまで、このゴルーチンは、他のゴルーチンをブロックします。
	c.v[key]--
	c.mu.Unlock()
}

// SafeCounterのメソッドであるValueは、カウンタの現在の値を返します。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// LockがUnlockされるまで、このゴルーチンは、他のゴルーチンをブロックします。
	defer c.mu.Unlock()
	fmt.Println("key:",c.v[key])
	return c.v[key]
}

