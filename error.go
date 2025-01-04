package atomicringbuffer

import "errors"

var (
	ErrIsEmpty = errors.New("buffer is empty")
	ErrIsFull  = errors.New("buffer is full")
)
