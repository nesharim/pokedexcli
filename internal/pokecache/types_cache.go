package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]CacheEntry
	mux   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
