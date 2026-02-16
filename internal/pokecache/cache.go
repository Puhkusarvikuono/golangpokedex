package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries	map[string]CacheEntry
	mu						sync.Mutex
}

type CacheEntry struct {
	createdAt		time.Time	
	val					[]byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		CacheEntries: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	//sync.Mutex to protect map
	c.mu.Lock()
	c.CacheEntries[key] = CacheEntry{
		createdAt: 	time.Now(), 
		val:				val,
	}
	c.mu.Unlock()
} 

func (c *Cache) Get(getKey string) ([]byte, bool) {
	c.mu.Lock()
	for cacheKey, cacheValue := range c.CacheEntries {
		if getKey == cacheKey {
			c.mu.Unlock()
			return cacheValue.val, true
		}
	}
	c.mu.Unlock()
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		start := time.Now()
		c.mu.Lock()
		for key, value := range c.CacheEntries {
			if start.Sub(value.createdAt) >= interval {
				delete(c.CacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
