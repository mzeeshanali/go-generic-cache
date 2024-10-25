package main

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
}

func main() {
	stringIntCache := NewCache[string, int](10 * time.Second)
	stringIntCache.Set("views", 100)
	if views, ok := stringIntCache.Get("views"); ok {
		fmt.Printf("Views: %d\n", views)
	}

	userCache := NewCache[string, User](5 * time.Second)
	userCache.Set("user1", User{ID: 1, Name: "John Doe"})
	if user, ok := userCache.Get("user1"); ok {
		fmt.Printf("User: %+v\n", user)
	}

	// create an LRU cache for string -> int with capacity 2
	lruCache := NewLRUCache[string, int](2)
	lruCache.Set("a", 1)
	lruCache.Set("b", 2)
	if value, ok := lruCache.Get("a"); ok {
		fmt.Printf("LRU Cache: a = %d\n", value)
	}

	lruCache.Set("c", 3)
	if _, ok := lruCache.Get("b"); !ok {
		fmt.Println("LRU Cache: b evicted")
	}
}
