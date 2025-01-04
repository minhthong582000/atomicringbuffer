# Ringbuffer

A lock-free single producer, single consumer ringbuffer implementation in Go using atomic operations.

## Installation

```bash
go get github.com/minhthong582000/atomicringbuffer
```

## Usage

```go
import "github.com/minhthong582000/atomicringbuffer"
```

Construct a new ringbuffer with a given size and type `T`:

```go
rb := atomicringbuffer.NewRingBuffer[int](1024) // capacity = 1024, type = int
```

Push an item to the back of the ringbuffer:

```go
err := rb.PushBack(2)
if err != nil {
    // Handle error
}
```

Pop an item from the front of the ringbuffer:

```go
item, err := rb.PopFront()
if err != nil {
    // Handle error
}

fmt.Println(item)
```

## Simple Example

A simple example of lock-free single producer, single consumer:

```bash
go run example/spsc/main.go
```

## License

Distributed under the GPLv3 License. See `LICENSE.md` file for more information.
