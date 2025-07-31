package pokecache

import(
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	values	map[string]cacheEntry
	mux		sync.Mutex
}

func NewCache(ttl time.Duration) *Cache{
	cache := &Cache{}
	cache.values = make(map[string]cacheEntry)
	go cache.reapLoop(ttl)
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	c.values[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.values[key]
	c.mux.Unlock()

	if ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		// Every time the ticker sends a Time
		for tick := range ticker.C {
			// Iterate through the map
			for k, v := range c.values {
				// Calculate the Time of expiration and compare it with the Time sent by the ticker
				expiration := v.createdAt.Add(interval)
				if tick.Compare(expiration) >= 0 {
					c.mux.Lock()
					delete(c.values, k)
					c.mux.Unlock()
				}
			}
		}
	}
}