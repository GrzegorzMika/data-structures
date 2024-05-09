package binaryheap

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func TestNewBinaryHeap(t *testing.T) {
	bh := NewBinaryHeap[int]()
	if bh.Len() != 0 {
		t.Errorf("Len() = %d, want 0", bh.Len())
	}
	if !bh.IsEmpty() {
		t.Errorf("IsEmpty() = %t, want true", bh.IsEmpty())
	}
	if bh.Cap() != 0 {
		t.Errorf("Cap() = %d, want 0", bh.Cap())
	}
}

func TestNewBinaryHeapWithCapacity(t *testing.T) {
	bh := NewBinaryHeapWithCapacity[int](10)
	if bh.Len() != 0 {
		t.Errorf("Len() = %d, want 0", bh.Len())
	}
	if !bh.IsEmpty() {
		t.Errorf("IsEmpty() = %t, want true", bh.IsEmpty())
	}
	if bh.Cap() != 10 {
		t.Errorf("Cap() = %d, want 10", bh.Cap())
	}
	bh.Clip()
	if bh.Len() != 0 {
		t.Errorf("Len() = %d, want 0", bh.Len())
	}
}

func TestBinaryHeapPush(t *testing.T) {
	type testCase struct {
		name             string
		elements         []int
		expectedOrdering []int
	}

	testCases := []testCase{
		{
			name:             "empty",
			elements:         []int{},
			expectedOrdering: []int{},
		},
		{
			name:             "one element",
			elements:         []int{1},
			expectedOrdering: []int{1},
		},
		{
			name:             "two elements",
			elements:         []int{1, 2},
			expectedOrdering: []int{2, 1},
		},
		{
			name:             "multiple elements",
			elements:         []int{17, 50, 32, 93, 8, 9, 69, 4, 26, 19, 16, 55, 6},
			expectedOrdering: []int{93, 50, 69, 26, 19, 55, 32, 4, 17, 8, 16, 9, 6},
		},
		{
			name:             "multiple elements with duplicates",
			elements:         []int{50, 71, 46, 78, 98, 54, 13, 67, 21, 3, 100, 91, 13, 54, 31, 28, 33, 30, 52, 68, 31, 71},
			expectedOrdering: []int{100, 98, 91, 67, 78, 54, 54, 50, 52, 68, 71, 46, 13, 13, 31, 28, 33, 21, 30, 3, 31, 71},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bh := NewBinaryHeapWithCapacity[int](len(tc.elements))
			bh.Push(tc.elements...)
			if bh.Len() != len(tc.expectedOrdering) {
				t.Errorf("Len() = %d, want %d", bh.Len(), len(tc.expectedOrdering))
			}
			for i := 0; i < bh.Len(); i++ {
				if bh.items[i] != tc.expectedOrdering[i] {
					t.Errorf("items[%d] = %d, want %d", i, bh.items[i], tc.expectedOrdering[i])
				}
			}
		})
	}

}

func TestBinaryHeapPeek(t *testing.T) {
	bh := NewBinaryHeapWithCapacity[int](10)
	x, ok := bh.Peek()
	if ok || x != 0 {
		t.Errorf("Peek() = (%v, %t), want (0, false)", x, ok)
	}
	bh.Push(17, 50, 32, 93, 8, 9, 69, 4, 26, 19)
	x, ok = bh.Peek()
	if !ok || x != 93 {
		t.Errorf("Pop() = (%v, %t), want (%d, true)", x, ok, 93)
	}
	if bh.Len() != 10 {
		t.Errorf("Len() = %d, want 10", bh.Len())
	}
	if bh.Cap() != 10 {
		t.Errorf("Cap() = %d, want 10", bh.Cap())
	}
}

func TestBinaryHeapRemove(t *testing.T) {
	bh := NewBinaryHeapWithCapacity[int](10)
	x, ok := bh.Pop()
	if ok || x != 0 {
		t.Errorf("Pop() = (%v, %t), want (0, false)", x, ok)
	}
	bh.Push(17, 50, 32, 93, 8, 9, 69, 4, 26, 19)
	x, ok = bh.Pop()
	if !ok || x != 93 {
		t.Errorf("Pop() = (%v, %t), want (%d, true)", x, ok, 93)
	}
	if bh.Len() != 9 {
		t.Errorf("Len() = %d, want 9", bh.Len())
	}
	if bh.Cap() != 10 {
		t.Errorf("Cap() = %d, want 10", bh.Cap())
	}
}

func TestBinaryHeapMultipleRemove(t *testing.T) {
	elements := []int{17, 50, 32, 93, 8, 9, 69, 4, 26, 19}
	bh := NewBinaryHeapWithCapacity[int](len(elements))
	bh.Push(elements...)
	slices.Sort(elements)
	slices.Reverse(elements)
	for _, v := range elements {
		log.Println(v)
		x, ok := bh.Pop()
		if !ok || x != v {
			t.Errorf("Pop() = (%v, %t), want (%d, true)", x, ok, v)
		}
	}
}

func BenchmarkBinaryHeapPush(b *testing.B) {
	bh := NewBinaryHeapWithCapacity[int](b.N)
	for i := range b.N {
		bh.Push(i)
	}
}

func ExampleBinaryHeap() {
	bh := NewBinaryHeap[int]()
	bh.Push(17, 50, 32, 93, 8, 9, 69, 4, 26, 19)
	fmt.Println(bh.Peek())
	v, ok := bh.Pop()
	if ok {
		fmt.Println(v)
	}
	// Output:
	// 93
	// 93
}
