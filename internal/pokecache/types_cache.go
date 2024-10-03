package pokecache

import (
	"time"
)

type Cache struct {
	entry map[string]CacheEntry
	// mu     sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
