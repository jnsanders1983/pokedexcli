package pokecache

import (
	"sync"
	"time"
)

// cacheEntry holds the raw data and its creation timestamp
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache wraps the map and a mutex to secure concurrent read/write access
type Cache struct {
	mu       sync.Mutex
	entry    map[string]cacheEntry
	interval time.Duration
}

// NewCache initializes and returns a new Cache instance
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		currentTime := time.Now()

		for key, entry := range c.entry {
			if currentTime.Sub(entry.createdAt) > c.interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
