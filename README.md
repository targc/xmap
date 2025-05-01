# xmap

**xmap** is a simple, lightweight, and efficient in-memory key-value store written in Go. Designed for speed and concurrency, it offers thread-safe operations and is ideal for building caching systems or temporary data stores.

## 🚀 Features

- ⚡ **High Performance** — Fast lookups, inserts, and deletes
- 🔒 **[TODO] Thread Safe** — Built-in concurrency protection using `sync.RWMutex`
- 🧠 **Lightweight** — Minimalistic design with zero external dependencies
- ⏱️ **[TODO] TTL Support** *(optional)* — Built-in expiration handling for keys
- 🧹 **Easy to Extend** — Clean and modular design

## 📦 Installation

```bash
go get github.com/targc/xmap
```

## 🛠️ Usage
```go
package main

import (
    "fmt"
    "github.com/targc/xmap"
)

func main() {
    store := xmap.New()

    store.Set("foo", "bar")
    value, ok := store.Get("foo")

    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }

    store.Delete("foo")
}
```
