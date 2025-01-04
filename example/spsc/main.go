package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/minhthong582000/atomicringbuffer"
)

var wg sync.WaitGroup

func Producer(rb *atomicringbuffer.RingBuffer[int]) {
	defer wg.Done()

	for i := 0; i < 20; i++ {
		_ = rb.PushBack(i)
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
	}
}

func Consumer(rb *atomicringbuffer.RingBuffer[int]) {
	defer wg.Done()

	for i := 0; i < 20; i++ {
		value, err := rb.PopFront()
		if err == nil {
			fmt.Println(value)
		}
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
	}
}

func main() {
	rb := atomicringbuffer.NewRingBuffer[int](10) // capacity = 10, T = int
	wg.Add(2)

	go Producer(rb)
	go Consumer(rb)

	wg.Wait()
}
