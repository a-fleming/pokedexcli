package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries:  make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}

	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exits := c.entries[key]
	if !exits {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-c.interval)

		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(cutoff) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
