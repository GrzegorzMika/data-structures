package binaryheap

import "cmp"

type BinaryHeap[T cmp.Ordered] struct {
	items []T
}

func NewBinaryHeap[T cmp.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{
		items: make([]T, 0),
	}
}

func NewBinaryHeapWithCapacity[T cmp.Ordered](capacity int) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		items: make([]T, 0, capacity),
	}
}

func (bh *BinaryHeap[T]) Len() int {
	return len(bh.items)
}

func (bh *BinaryHeap[T]) IsEmpty() bool {
	return len(bh.items) == 0
}

func (bh *BinaryHeap[T]) Cap() int {
	return cap(bh.items)
}

func (bh *BinaryHeap[T]) Push(xs ...T) {
	for _, x := range xs {
		bh.push(x)
	}
}

func (bh *BinaryHeap[T]) push(x T) {
	n := bh.Len()
	bh.items = append(bh.items, x)
	bh.bubbleUp(n)
}

func (bh *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (bh *BinaryHeap[T]) bubbleUp(i int) {
	parentIndex := bh.parent(i)
	for i > 0 && bh.items[i] >= bh.items[parentIndex] {
		bh.items[i], bh.items[parentIndex] = bh.items[parentIndex], bh.items[i]
		i = parentIndex
		parentIndex = bh.parent(i)
	}
}
