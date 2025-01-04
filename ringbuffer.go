package atomicringbuffer

import (
	"fmt"
	"sync/atomic"
)

// RingBuffer
type RingBuffer[T any] struct {
	capacity   uint64
	start, end atomic.Uint64
	buffer     []T
}

func NewRingBuffer[T any](capacity uint64) *RingBuffer[T] {
	return &RingBuffer[T]{
		capacity: capacity + 1,
		buffer:   make([]T, capacity+1),
	}
}

func (r *RingBuffer[T]) IsFull() bool {
	return r.incrementIndex(r.end.Load()) == r.start.Load()
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.start.Load() == r.end.Load()
}

func (r *RingBuffer[T]) Capacity() uint64 {
	return r.capacity - 1
}

func (r *RingBuffer[T]) Size() uint64 {
	return r.end.Load() - r.start.Load()
}

func (r *RingBuffer[T]) StartIndex() uint64 {
	return r.start.Load()
}

// incrementIndex use modulus to calculate when the index should wrap to the beginning
// in a circular way
func (r *RingBuffer[T]) incrementIndex(index uint64) uint64 {
	return (index + 1) % r.capacity
}

func (r *RingBuffer[T]) PushBack(value T) error {
	currEnd := r.end.Load()

	newEnd := r.incrementIndex(currEnd)
	if newEnd == r.start.Load() {
		return ErrIsFull
	}

	r.buffer[currEnd] = value
	r.end.Store(newEnd)

	return nil
}

func (r *RingBuffer[T]) PopFront() (T, error) {
	currStart := r.start.Load()

	if currStart == r.end.Load() {
		return *new(T), ErrIsEmpty
	}

	value := r.buffer[currStart]
	r.start.Store(r.incrementIndex(currStart))
	fmt.Println(r.buffer)
	return value, nil
}

func (r *RingBuffer[T]) PeekFront() (T, error) {
	start := r.start.Load()

	if start == r.end.Load() {
		return *new(T), ErrIsEmpty
	}

	value := r.buffer[start]

	return value, nil
}
