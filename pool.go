// Copyright 2022 Steven Dake
//
// ASL2
// This is a bitmapped object cache offering O(1) ammortized time focused on space savings.
//
// Goals:
// - O(1) ammortized time
// - Trade off time for space
//
// Non-goals:
// - deterministic performance

package pool

import (
	"math/bits"
	"reflect"
)

type Handle struct {
	line  uint16
	entry uint16
}

type Object struct {
	a int
	b int
	c int
}

type poolSlice[T Object] [][]T

type Pool[T Object] struct {
	// Internal Mutex of the
	//mu Mutex

	// Number of lines
	lines uint16

	// Lengths of line
	lineLengths []uint16

	// Size of object
	objectSize uint16

	objects poolSlice[T]

	bitmap []uint64
}

// Put sets the object x at line in O(1) ammortized time
func (p Pool[T]) Put(value T, line uint16) Handle {
	zeros := bits.TrailingZeros64(p.bitmap[line])
	if zeros == 0 {
		return Handle{entry: 0, line: 0}
	}

	zeros = zeros - 1

	p.objects[line][zeros] = value
	p.bitmap[line] |= 1 << zeros
	return Handle{entry: uint16(zeros), line: line}
}

// Get returns the object mapped to the Handle in O(1) time
// If no object is found, nil is returned.
func (p Pool[T]) Get(handle Handle) (value any) {
	return p.objects[handle.line][handle.entry]
}

// Remove eliminates the object referenced by Handle in O(1) ammortized time
func (p Pool[T]) Remove(handle Handle) {
	p.bitmap[uint64(handle.line)] &^= 1 << handle.entry
}

// Size returns a rough estimate of the number of bytes used by the pool.
func (p Pool[T]) Size() uint16 {
	var entries uint16 = 0
	var i uint16 = 0

	for i = 0; i < p.lines; i++ {
		entries += p.lineLengths[i]
	}
	return p.objectSize * entries
}

// New returns a new steak pool with objects of type object that includes
// an object cache with count of lines. Each line size is passed by a variadic.
// with a count of lines. Each line size is represented by a variadic.
// All inputs must be a power of two. An error returns nil in .
func New[T Object](object T, lines uint16, lineLengths ...uint16) *Pool[T] {
	pool := new(Pool[T])
	pool.lineLengths = make([]uint16, lines, lines)
	pool.objects = make([][]T, lines, lines)
	pool.bitmap = make([]uint64, lines, lines)
	pool.objectSize = uint16(reflect.TypeOf(object).Size())
	pool.lines = lines
	pool.lineLengths = lineLengths

	var i uint16 = 0
	for i = 0; i < lines; i++ {
		pool.objects[i] = make([]T, lineLengths[i], lineLengths[i])
		pool.bitmap[i] = 0
	}

	return pool
}
