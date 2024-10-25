# Generic Cache Implementation in Go

This project provides a generic in-memory cache implementation in Go, including both a time-based cache and a Least Recently Used (LRU) cache. The implementation leverages Go's generics feature, allowing you to create caches for various key-value types efficiently.

## Features

- **Time-Based Cache:** Stores items with a specified expiration time.
- **LRU Cache:** Implements an eviction policy that removes the least recently used items when the cache reaches its capacity.
- **Thread-Safe:** Both cache implementations are thread-safe and can be used concurrently.

## Installation

Ensure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

1. Clone this repository:

   ```bash
   git clone https://github.com/mzeeshanali/go-generic-cache.git
   cd go-generic-cache
   ```
2. To Run:
   
   ```bash
   go run .
   ```
## Usage
### Time-Based Cache
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a string -> int cache with a 10-second lifetime
	stringIntCache := NewCache[string, int](10 * time.Second)
	stringIntCache.Set("views", 100)

	if views, ok := stringIntCache.Get("views"); ok {
		fmt.Printf("Views: %d\n", views)
	}
}
```

### LRU Cache
```go
package main

import (
	"fmt"
)

func main() {
	// Create an LRU cache for string -> int with a capacity of 2
	lruCache := NewLRUCache 
	lruCache.Set("a", 1)
	lruCache.Set("b", 2)

	if value, ok := lruCache.Get("a"); ok {
		fmt.Printf("LRU Cache: a = %d\n", value)
	}

	lruCache.Set("c", 3) // This will evict "b" since it's least recently used
	if _, ok := lruCache.Get("b"); !ok {
		fmt.Println("LRU Cache: b evicted")
	}
}
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
