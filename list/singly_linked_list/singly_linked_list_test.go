package singlylinkedlist

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	l := NewSinglyLinkedList[int]()

	if l.head != nil {
		t.Errorf("expected head to be nil, got %#v", l.head)
	}
	if l.Len() != 0 {
		t.Errorf("expected length to be 0, got %d", l.length)
	}
}

func TestPushFront(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.PushFront(1)
	l.PushFront(2)

	if l.Len() != 2 {
		t.Errorf("expected length to be 2, got %d", l.length)
	}
	if *l.head.Data != 2 {
		t.Errorf("expected head to be 2, got %d", *l.head.Data)
	}
	if *l.head.next.Data != 1 {
		t.Errorf("expected head.next to be 1, got %d", *l.head.next.Data)
	}
}

func TestPushBack(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)

	if l.Len() != 2 {
		t.Errorf("expected length to be 2, got %d", l.length)
	}
	if *l.head.Data != 1 {
		t.Errorf("expected head to be 1, got %d", *l.head.Data)
	}
	if *l.head.next.Data != 2 {
		t.Errorf("expected head.next to be 1, got %d", *l.head.next.Data)
	}
}

func TestPopFront(t *testing.T) {
	l := NewSinglyLinkedList[int]()

	x, ok := l.PopFront()
	if ok || x != 0 {
		t.Errorf("PopFront() = (%v, %t), want (0, false)", x, ok)
	}

	l.PushFront(1)
	l.PushFront(2)

	x, ok = l.PopFront()
	if !ok || x != 2 {
		t.Errorf("PopFront() = (%v, %t), want (2, true)", x, ok)
	}
	if l.Len() != 1 {
		t.Errorf("expected length to be 1, got %d", l.length)
	}
	x, ok = l.PopFront()
	if !ok || x != 1 {
		t.Errorf("PopFront() = (%v, %t), want (1, true)", x, ok)
	}
	if l.Len() != 0 {
		t.Errorf("expected length to be 0, got %d", l.length)
	}
	x, ok = l.PopFront()
	if ok || x != 0 {
		t.Errorf("PopFront() = (%v, %t), want (0, false)", x, ok)
	}
}

func TestPopBack(t *testing.T) {
	l := NewSinglyLinkedList[int]()

	x, ok := l.PopBack()
	if ok || x != 0 {
		t.Errorf("PopBack() = (%v, %t), want (0, false)", x, ok)
	}

	l.PushBack(1)
	l.PushBack(2)

	x, ok = l.PopBack()
	if !ok || x != 2 {
		t.Errorf("PopBack() = (%v, %t), want (2, true)", x, ok)
	}
	if l.Len() != 1 {
		t.Errorf("expected length to be 1, got %d", l.length)
	}
	x, ok = l.PopBack()
	if !ok || x != 1 {
		t.Errorf("PopBack() = (%v, %t), want (1, true)", x, ok)
	}
	if l.Len() != 0 {
		t.Errorf("expected length to be 0, got %d", l.length)
	}
	x, ok = l.PopBack()
	if ok || x != 0 {
		t.Errorf("PopBack() = (%v, %t), want (0, false)", x, ok)
	}
}

func TestContains(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	cmp := func(x1, x2 int) bool { return x1 == x2 }

	if l.Contains(1, cmp) {
		t.Errorf("Contains(1) = true, want false")
	}

	l.PushBack(1)
	l.PushBack(2)

	if !l.Contains(1, cmp) {
		t.Errorf("Contains(1) = false, want true")
	}
	if !l.Contains(2, cmp) {
		t.Errorf("Contains(2) = false, want true")
	}
	if l.Contains(3, cmp) {
		t.Errorf("Contains(3) = true, want false")
	}
}

func TestFront(t *testing.T) {
	l := NewSinglyLinkedList[int]()

	x, ok := l.Front()
	if ok || x != nil {
		t.Errorf("Front() = (%v, %t), want (nil, false)", x, ok)
	}

	l.PushBack(1)
	l.PushBack(2)

	x, ok = l.Front()
	if !ok || *x.Data != 1 {
		t.Errorf("Front() = (%v, %t), want (1, true)", *x.Data, ok)
	}

	_, _ = l.PopFront()
	x, ok = l.Front()
	if !ok || *x.Data != 2 {
		t.Errorf("Front() = (%v, %t), want (2, true)", *x.Data, ok)
	}
}

func TestBack(t *testing.T) {
	l := NewSinglyLinkedList[int]()

	x, ok := l.Back()
	if ok || x != nil {
		t.Errorf("Back() = (%v, %t), want (nil, false)", x, ok)
	}

	l.PushBack(1)
	l.PushBack(2)

	x, ok = l.Back()
	if !ok || *x.Data != 2 {
		t.Errorf("Back() = (%v, %t), want (2, true)", *x.Data, ok)
	}

	_, _ = l.PopBack()
	x, ok = l.Back()
	if !ok || *x.Data != 1 {
		t.Errorf("Back() = (%v, %t), want (1, true)", *x.Data, ok)
	}
}

func TestClear(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)

	l.Clear()
	if l.Len() != 0 {
		t.Errorf("expected length to be 0, got %d", l.length)
	}
	if l.head != nil {
		t.Errorf("expected head to be nil, got %#v", l.head)
	}
}

func TestRemove(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	x := l.Remove(2)
	if x != 3 {
		t.Errorf("Remove(2) = %d, want 3", x)
	}
	if l.Len() != 2 {
		t.Errorf("expected length to be 2, got %d", l.length)
	}
	x = l.Remove(0)
	if x != 1 {
		t.Errorf("Remove(0) = %d, want 1", x)
	}
}

func TestRemovePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	l := NewSinglyLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(4)
}

func TestAppend(t *testing.T) {
	type testCase struct {
		name        string
		prepareFunc func() (*SinglyLinkedList[int], *SinglyLinkedList[int])
		checkFunc   func(*testing.T, *SinglyLinkedList[int])
	}

	testCases := []testCase{
		{
			name: "empty",
			prepareFunc: func() (*SinglyLinkedList[int], *SinglyLinkedList[int]) {
				l1 := NewSinglyLinkedList[int]()
				l2 := NewSinglyLinkedList[int]()
				return l1, l2
			},
			checkFunc: func(t *testing.T, l1 *SinglyLinkedList[int]) {
				if l1.Len() != 0 {
					t.Errorf("expected length to be 0, got %d", l1.length)
				}
				if l1.head != nil {
					t.Errorf("expected head to be nil, got %#v", l1.head)
				}
			},
		},
		{
			name: "append empty to non-empty",
			prepareFunc: func() (*SinglyLinkedList[int], *SinglyLinkedList[int]) {
				l1 := NewSinglyLinkedList[int]()
				l2 := NewSinglyLinkedList[int]()
				l1.PushBack(1)
				return l1, l2
			},
			checkFunc: func(t *testing.T, l1 *SinglyLinkedList[int]) {
				if l1.Len() != 1 {
					t.Errorf("expected length to be 0, got %d", l1.length)
				}
				if l1.head.next != nil {
					t.Errorf("expected head.next to be nil, got %#v", l1.head)
				}
			},
		},
		{
			name: "append non-empty to empty",
			prepareFunc: func() (*SinglyLinkedList[int], *SinglyLinkedList[int]) {
				l1 := NewSinglyLinkedList[int]()
				l2 := NewSinglyLinkedList[int]()
				l2.PushBack(1)
				return l1, l2
			},
			checkFunc: func(t *testing.T, l1 *SinglyLinkedList[int]) {
				if l1.Len() != 1 {
					t.Errorf("expected length to be 1, got %d", l1.length)
				}
				if l1.head == nil {
					t.Errorf("expected head to be non-nil, got %#v", l1.head)
				}
			},
		},
		{
			name: "append non-empty to non-empty",
			prepareFunc: func() (*SinglyLinkedList[int], *SinglyLinkedList[int]) {
				l1 := NewSinglyLinkedList[int]()
				l2 := NewSinglyLinkedList[int]()
				l1.PushBack(1)
				l1.PushBack(2)
				l2.PushBack(3)
				l2.PushBack(4)
				return l1, l2
			},
			checkFunc: func(t *testing.T, l1 *SinglyLinkedList[int]) {
				if l1.Len() != 4 {
					t.Errorf("expected length to be 4, got %d", l1.length)
				}
				if *l1.head.Data != 1 {
					t.Errorf("expected head to be 1, got %d", *l1.head.Data)
				}

				x, ok := l1.Back()
				if !ok || *x.Data != 4 {
					t.Errorf("Back() = (%v, %t), want (4, true)", *x.Data, ok)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1, l2 := tc.prepareFunc()
			l1.Append(l2)
			tc.checkFunc(t, l1)
		})
	}
}
