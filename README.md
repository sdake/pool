# Pool
[![Go Report Card](https://goreportcard.com/badge/github.com/sdake/pool)](https://goreportcard.com/report/github.com/sdake/pool)

✨ **`pool` is a bitmapped object cache implemented implemented  with Go generic support.**

## ❓ Why

- I wanted to learn Go generic syntax.
- I wanted to minimize the memory consumption of object storage.
- I wanted objects addressible by lines as well as handles.

## 💡 Usage

You can import `pool` using:

```go
import (
	"github.com/sdake/pool"
)
```

Then create an object type, identify the number and length of the cache lines, and then create pool:

```go
// Define your custom object that will be stored in the cache lines.
type Object {
	abra int
	cadbra int
	popcorn string
}

// Define the cache lines sizes. The minimum size is 64 entries. Cache line sizes must be multiples of 64.
lineSize := []uint16{64, 64, 64, 64}

// Create a pool object with 4 cache lines
pool := pool.New(Object, 4, lineSize...)
```

Initialize an object and add to the pool in the third cache line:
```
ObjectS := Object{2, 1, "with butter please!"}
handle := pool.Put(ObjectS, 3)
```

Remove that same object:
```
pool.Remove(handle)
```

## 👤 Authors

- Steven Dake

## 💫 Show your support

I don't accept donations. So please give a ⭐️ if this project helped you!


## 📝 License

Copyright © 2022 [Steven Dake](https://github.com/sdake/).

This project is [ASL2](./LICENSE) licensed.
