// Package binaryheap implements the generic binary heap.
package binaryheap

import (
	"cmp"
	"slices"
)

type BinaryHeap[T cmp.Ordered] struct {
	items []T
}

// NewBinaryHeap creates a new instance of BinaryHeap.
// The binary heap is a complete binary tree where each node is bigger or equal to its children.
// The elements are stored in an array, and the heap property is maintained with every new element added to the heap.
func NewBinaryHeap[T cmp.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{
		items: make([]T, 0),
	}
}

// NewBinaryHeapWithCapacity creates a new instance of BinaryHeap with the specified capacity.
// The capacity is the maximum number of elements that the binary heap can hold without reallocating its underlying slice.
func NewBinaryHeapWithCapacity[T cmp.Ordered](capacity int) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		items: make([]T, 0, capacity),
	}
}

// Len returns the number of elements in the binary heap.
//
// The time complexity of this method is O(1).
func (bh *BinaryHeap[T]) Len() int {
	return len(bh.items)
}

// IsEmpty checks if the binary heap is empty.
//
// It returns true if the binary heap has no elements, and false otherwise.
//
// The time complexity of this method is O(1).
func (bh *BinaryHeap[T]) IsEmpty() bool {
	return len(bh.items) == 0
}

// Cap returns the capacity of the binary heap.
//
// The capacity is the maximum number of elements that the binary heap can hold
// without reallocating its underlying slice.
func (bh *BinaryHeap[T]) Cap() int {
	return cap(bh.items)
}

// Clip removes unused capacity from the binary heap.
// It is useful when the capacity of the binary heap is no longer needed and can be
// reclaimed by the garbage collector.
//
// Clip does not change the length of the binary heap; it merely resizes the capacity.
//
// It is safe to call Clip even if the binary heap is already at its capacity.
func (bh *BinaryHeap[T]) Clip() {
	bh.items = slices.Clip(bh.items)
}

// Push adds one or more elements to the binary heap.
//
// The Push method takes a variadic parameter xs, which represents the elements to be added to the heap.
//
// The time complexity of adding each element is O(log n), where n is the number of elements in the heap.
func (bh *BinaryHeap[T]) Push(xs ...T) {
	for _, x := range xs {
		bh.push(x)
	}
}

// Peek returns the biggest element in the binary heap without removing it and true.
// If the binary heap is empty, it returns a zero value of type T and false.
//
// The time complexity of this method is O(1).
func (bh *BinaryHeap[T]) Peek() (T, bool) {
	if bh.Len() == 0 {
		return *new(T), false
	}
	return bh.items[0], true
}

// Pop removes and returns the biggest element from the binary heap.
// If the binary heap is empty, it returns a zero value of type T and false.
//
// The time complexity of this method is O(log n), where n is the number of elements in the heap.
func (bh *BinaryHeap[T]) Pop() (T, bool) {
	if bh.Len() == 0 {
		return *new(T), false
	}
	x := bh.items[0]
	bh.items[0] = bh.items[bh.Len()-1]
	bh.items = bh.items[:bh.Len()-1]
	bh.sinkDown(0)
	return x, true
}

func (bh *BinaryHeap[T]) push(x T) {
	n := bh.Len()
	bh.items = append(bh.items, x)
	bh.bubbleUp(n)
}

func (bh *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (bh *BinaryHeap[T]) left(i int) int {
	return 2*i + 1
}
func (bh *BinaryHeap[T]) right(i int) int {
	return 2*i + 2
}

func (bh *BinaryHeap[T]) bubbleUp(i int) {
	parentIndex := bh.parent(i)
	for i > 0 && bh.items[i] >= bh.items[parentIndex] {
		bh.items[i], bh.items[parentIndex] = bh.items[parentIndex], bh.items[i]
		i = parentIndex
		parentIndex = bh.parent(i)
	}
}

func (bh *BinaryHeap[T]) singleStepDown(i int) int {
	j := -1
	r := bh.right(i)
	if r < bh.Len() && bh.items[r] >= bh.items[i] {
		l := bh.left(i)
		if bh.items[l] >= bh.items[r] {
			j = l
		} else {
			j = r
		}
	} else {
		l := bh.left(i)
		if l < bh.Len() && bh.items[l] >= bh.items[i] {
			j = l
		}
	}
	if j >= 0 {
		bh.items[i], bh.items[j] = bh.items[j], bh.items[i]
	}
	i = j
	return i
}

func (bh *BinaryHeap[T]) sinkDown(i int) {
	i = bh.singleStepDown(i)
	for i >= 0 {
		i = bh.singleStepDown(i)
	}
}
