package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}
type Cache struct {
	cache map[string]cacheEntry
	mut   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	var cache = Cache{
		cache: make(map[string]cacheEntry),
		mut:   sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Time{}.UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cache[key]
	if ok {
		return val.val, ok
	}
	return []byte{}, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		if c.mut.TryLock() {
			now := time.Now()
			for key, val := range c.cache {
				if val.createdAt.Before(now.Add(-interval * time.Second)) {
					delete(c.cache, key)
				}
			}
			c.mut.Unlock()
		}
	}
}
