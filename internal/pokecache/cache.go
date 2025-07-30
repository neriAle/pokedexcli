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
	// Call reapLoop()
	cache := &Cache{}
	return cache
}