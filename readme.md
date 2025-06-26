# go-ordered-id

A lightweight Go package to generate **globally ordered**, **URL-safe**, and **short unique IDs** using a combination of:
- Microsecond-precision timestamps (encoded in Base62)
- A short random suffix (NanoID-based)

Perfect for use cases like:
- Public short URLs
- Distributed job IDs
- Ordered message/event IDs
- Alternative to UUID/ULID for compact ID generation

---


## 🚀 Features

- ✅ Globally ordered (lexicographically by timestamp)
- ✅ URL-safe Base62 encoding
- ✅ Short IDs (e.g. 12–14 characters)
- ✅ High uniqueness (62⁴ = 14M+ per microsecond)
- ✅ No external dependencies (except `go-nanoid/v2`)
- ✅ Fast and thread-safe

This module is based on [go-nanoid v2](github.com/matoous/go-nanoid/v2)

---

## 📦 Installation

```bash
go get github.com/toni-ismail/go-ordered-id
```

---

## ✨ Usage

```go
package main

import (
	"fmt"
	"github.com/toni-ismail/go-ordered-id"
)

func main() {
	id := orderedid.Generate()
	fmt.Println("Generated ID:", id)
}
```

Example output:

```
6SzgrA4DU9g
6SzgrA4DU9h
```

> The prefix is a Base62-encoded `UnixMicro()`, followed by a 4-char random NanoID suffix.

---

## 🔒 Uniqueness

- Supports ~14.7 million combinations per microsecond (`62^4`)
- Can safely generate hundreds of thousands of IDs per second without collision
- Uses a lock + timestamp deduplication to ensure monotonic order

---

## 🔧 Customization (future work)

You can easily fork and adjust:
- Suffix length (e.g., 5–6 for more entropy)
- Use milliseconds instead of microseconds
- Add machine/shard ID prefix (for clustered setups)

---

## 🧪 Benchmark

```
BenchmarkGenerate-8   	500000	      2400 ns/op	    120 B/op	    2 allocs/op
```

> This shows you can generate ~400K+ ordered IDs/sec on a single core.

---

## 📜 License

MIT

---

## ✍️ Author

Made by [@yourusername](https://github.com/yourusername). Contributions welcome!
