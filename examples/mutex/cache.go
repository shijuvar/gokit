package main

import (
	"fmt"
	"sync"
)

type MemCache struct {
	// A RWMutex is a reader/writer mutual exclusion lock.
	// The lock can be held by an arbitrary number of readers or a single writer.
	sync.RWMutex
	cache map[string]string
}

func NewInMemoryCache() *MemCache {
	return &MemCache{
		cache: make(map[string]string),
	}
}
func (ds *MemCache) set(key string, value string) {
	ds.cache[key] = value
}
func (ds *MemCache) get(key string) string {
	if ds.count() > 0 {
		item := ds.cache[key]
		return item
	}
	return ""
}
func (ds *MemCache) count() int {
	return len(ds.cache)
}
func (ds *MemCache) Set(key string, value string) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}
func (ds *MemCache) Get(key string) string {
	ds.RLock()
	defer ds.RUnlock()
	return ds.get(key)
}
func (ds *MemCache) Count() int {
	ds.RLock()
	defer ds.RUnlock()
	return ds.count()
}
func main() {
	cache := NewInMemoryCache()
	cache.Set("dev", "Gopher")
	val := cache.Get("dev")
	fmt.Println(val)
}
