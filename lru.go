package main

import (
	"container/list"
	"sync"
)

// LRUCache is a generic cache with LRU eviction policy.
type LRUCache[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	items    map[K]*list.Element
	order    *list.List
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLRUCache creates a new LRUCache with the specified capacity.
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
		order:    list.New(),
	}
}

// Set adds or updates an item in the cache.
func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		el.Value.(*entry[K, V]).value = value
		return
	}

	if c.order.Len() >= c.capacity {
		c.evict()
	}

	ent := &entry[K, V]{key: key, value: value}
	el := c.order.PushFront(ent)
	c.items[key] = el
}

// Get retrieves an item from the cache and marks it as most recently used.
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		return el.Value.(*entry[K, V]).value, true
	}

	var zeroValue V
	return zeroValue, false
}

// evict removes the least recently used item from the cache.
func (c *LRUCache[K, V]) evict() {
	el := c.order.Back()
	if el == nil {
		return
	}

	c.order.Remove(el)
	ent := el.Value.(*entry[K, V])
	delete(c.items, ent.key)
}
