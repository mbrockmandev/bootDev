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
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
}

func (c *Cache) Get(key string) ([]byte, bool) {
}

func (c *Cache) reapLoop(interval time.Duration) {
}

func (c *Cache) reap(now time.Time, last time.Duration) {
}
