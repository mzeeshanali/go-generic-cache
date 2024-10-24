package main

import (
	"sync"
	"time"
)

// CacheItem represents a single item in the cache.
type CacheItem[V any] struct {
	Value      V
	Expiration time.Time
}

// Cache is a generic in-memory cache that supports any type of key (K) and value (V).
type Cache[K comparable, V any] struct {
	mu       sync.RWMutex
	items    map[K]CacheItem[V]
	lifetime time.Duration
}

// NewCache creates a new Cache with a specified item lifetime.
func NewCache[K comparable, V any](lifetime time.Duration) *Cache[K, V] {
	return &Cache[K, V]{
		items:    make(map[K]CacheItem[V]),
		lifetime: lifetime,
	}
}

// Set adds a new item to the cache or updates an existing item.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(c.lifetime)
	c.items[key] = CacheItem[V]{Value: value, Expiration: expiration}
}

// Get retrieves an item from the cache.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || item.Expiration.Before(time.Now()) {
		var zeroValue V
		return zeroValue, false
	}
	return item.Value, true
}

// Delete removes an item from the cache.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Size returns the current number of items in the cache.
func (c *Cache[K, V]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}
