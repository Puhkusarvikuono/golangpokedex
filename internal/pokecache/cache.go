package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cache struct {
	CacheEntry		map[string]CacheEntry
	mu						sync.Mutex
}

type CacheEntry struct {
	createdAt		time.Time	
	val					[]byte
}

func NewCache(interval int) (cache, error) {
	// calls cache.reapLoop()
	// so maybe pass interval for time.Duration
	newCache, err := cache{
		CacheEntry: //something
		mu: //something
	}
	if err != nil {
		return nil, err
	}
	return newCache, nil
}

func (c cache) cache.Add(key string, val []byte) () {
	//sync.Mutex to protect map
	//adds a new entry
} 

func (c cache) cache.Get(key string) ([]byte, bool) {
	//sync.Mutex to protect map
	// true if entry is found false if not	
	// check if "key" in cache
	// Return cache entry
}

func (c cache) cache.reapLoop() {
	//sync.Mutex to protect map
	//each time an interval (time.Duration newCache) passes it should remove any entries
	//that are older than the interval
	// If interval = 5s, entry 7s old will be removed
	// time.Ticker perhaps see tips
} 

