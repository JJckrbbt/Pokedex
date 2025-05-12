package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.RWMutex
	entries  map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go reapLoop(interval, cache)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result, exists := c.entries[key]
	if exists {
		return result.val, true
	} else {
		return nil, false
	}
}

func reapLoop(interval time.Duration, cache *Cache) {
	for {
		time.Sleep(cache.interval)
		cache.mu.Lock()

		deleteList := []string{}
		for name, entry := range cache.entries {
			age := time.Since(entry.createdAt)
			if age > interval {
				deleteList = append(deleteList, name)
			}
		}

		for _, entry := range deleteList {
			delete(cache.entries, entry)
		}
		cache.mu.Unlock()

	}
}
