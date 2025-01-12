package atomicringbuffer

import (
	"sync"
	"testing"
)

func BenchmarkRingBuffer(b *testing.B) {
	r := NewRingBuffer[int](uint64(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PushBack(i)
		_, _ = r.PopFront()
	}
}

func BenchmarkRingBufferSingleProducerSingleConsumer(b *testing.B) {
	var wg sync.WaitGroup
	r := NewRingBuffer[int](uint64(b.N))
	b.ResetTimer()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < b.N; i++ {
			_ = r.PushBack(i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < b.N; i++ {
			_, _ = r.PopFront()
		}
	}()

	wg.Wait()
}
