package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
