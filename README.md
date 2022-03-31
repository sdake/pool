# Pool
[![Go report](https://goreportcard.com/badge/github.com/sdake/pool)](https://goreportcard.com/report/github.com/sdake/pool)

✨ **`pool` is a bitmapped object cache implemented using Go generics.**

## 💡 Why

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

	// Define the cache lines sizes. The minimum size is 64 entries.
        lineSize := []uint16{64, 64, 64, 64}

	// Create a pool object with 4 cache lines
        pool := pool.New(Object, 4, lineSize...)
```

Add an object to the pool in the third cache line:
```
	handle = pool.Put(ObjectQ, 3)
```

Remove that same object:
```
	handle = pool.Remove(handle)
```

## 👤 Authors

- Steven Dake

## 💫 Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2022 [Steven Dake](https://github.com/sdake).

This project is [ASL2](./LICENSE) licensed.
