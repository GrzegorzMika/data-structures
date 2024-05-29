// Package singlylinkedlist implements the generic singly linked list.
package singlylinkedlist

import "fmt"

type Node[T any] struct {
	Data *T
	next *Node[T]
}

type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

// NewSinglyLinkedList creates and returns a new instance of a singly linked list.
// The type parameter T is used to specify the type of elements stored in the list.
// The function returns a pointer to a new SinglyLinkedList instance.
//
// Example usage:
//
//	list := singlylinkedlist.NewSinglyLinkedList[int]()
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Append appends all elements from the other list to the end of the current list.
// If the other list is empty, this function does nothing.
// Time complexity: O(n), where n is the length of the other list.
func (l *SinglyLinkedList[T]) Append(other *SinglyLinkedList[T]) {
	if other.Len() == 0 {
		return
	}
	if l.length == 0 {
		l.head = other.head
		l.length = other.length
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = other.head
	l.length += other.length
}

// Len returns the number of elements in the list.
// The time complexity of this operation is O(1).
func (l *SinglyLinkedList[T]) Len() int {
	return l.length
}

// Clear removes all elements from the list.
// It iterates through the list, setting the next pointer of each node to nil,
// allowing the garbage collector to reclaim the memory occupied by the nodes.
// After clearing the list, it sets the head and length to nil and 0 respectively.
// The time complexity of this operation is O(n), where n is the length of the list.
func (l *SinglyLinkedList[T]) Clear() {
	// dereference object to allow GC to clean up
	for n := l.head; n != nil; {
		tmp := n.next
		n.next = nil
		n = tmp
	}
	l.length = 0
	l.head = nil
}

// Contains checks if the list contains a specific element based on the provided equality function.
// It iterates through the list and compares each element using the provided equality function.
// If an element is found that satisfies the equality function, it returns true.
// If no such element is found, it returns false.
func (l *SinglyLinkedList[T]) Contains(x T, eq func(T, T) bool) bool {
	for n := l.head; n != nil; n = n.next {
		if eq(*n.Data, x) {
			return true
		}
	}
	return false
}

// Front returns the first node of the list and a boolean indicating if the list is not empty.
// If the list is empty, it returns nil and false.
//
// The function checks if the length of the list is zero. If it is, it returns nil and false.
// Otherwise, it returns a pointer to the head node and true.
//
// The time complexity of this operation is O(1), as it only requires accessing the head pointer.
func (l *SinglyLinkedList[T]) Front() (*Node[T], bool) {
	if l.length == 0 {
		return nil, false
	}
	return l.head, true
}

// Back returns the last node of the list and a boolean indicating if the list is not empty.
// If the list is empty, it returns nil and false.
//
// The function iterates through the list until it finds the last node (i.e., a node where next is nil).
// It then returns a pointer to the last node and true.
// If the list is empty (i.e., head is nil), it returns nil and false.
//
// The time complexity of this operation is O(n), where n is the length of the list.
func (l *SinglyLinkedList[T]) Back() (*Node[T], bool) {
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			return n, true
		}
	}
	return nil, false
}

// Remove removes and returns the element at the specified index from the list.
// If the index is out of range, it panics with an error message.
// If the index is 0, it updates the head of the list to the next node.
// Otherwise, it traverses the list to find the node at the specified index,
// updates the previous node's next pointer to skip the current node,
// and decrements the list's length.
// The time complexity of this operation is O(n), where n is the length of the list.
func (l *SinglyLinkedList[T]) Remove(at int) T {
	if at < 0 || at > l.length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", at, l.length))
	}
	var x T
	// handle a special case where we need to update the head of the list
	if at == 0 {
		x = *l.head.Data
		l.head = l.head.next
		l.length--
		return x
	}
	previous := l.head
	for i := 1; i < at; i++ {
		previous = previous.next
	}
	x = *previous.next.Data
	previous.next = previous.next.next
	l.length--
	return x
}

// PopFront removes and returns the first element from the list.
// If the list is empty, it returns the zero value of type T and false.
// Otherwise, it updates the head of the list to the next node, decrements the list's length,
// and returns the value of the removed node.
// The time complexity of this operation is O(1), as it only requires updating the head pointer.
func (l *SinglyLinkedList[T]) PopFront() (T, bool) {
	if l.length == 0 {
		return *new(T), false
	}
	x := *l.head.Data
	l.head = l.head.next
	l.length--
	return x, true
}

// PopBack removes and returns the last element from the list.
// If the list is empty, it returns the zero value of type T and false.
// Otherwise, it updates the last node's next pointer to nil, decrements the list's length,
// and returns the value of the removed node.
// The time complexity of this operation is O(n), where n is the length of the list.
func (l *SinglyLinkedList[T]) PopBack() (T, bool) {
	if l.length == 0 {
		return *new(T), false
	}
	var x T
	previous := l.head
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			x = *n.Data
			previous.next = nil
			l.length--
			break
		}
		previous = n
	}
	return x, true
}

// PushFront adds a new node with the given value to the front of the list.
// It creates a new node with the provided value and sets its next pointer to the current head of the list.
// Then, it updates the head of the list to the newly created node.
// The time complexity of this operation is O(1), as it only requires updating the head pointer.
func (l *SinglyLinkedList[T]) PushFront(x T) {
	// Create a new node with the given value and the current head as its next pointer
	node := &Node[T]{
		Data: &x,
		next: l.head,
	}
	// Update the head of the list to the newly created node
	l.head = node
	// Increment the length of the list
	l.length++
}

// PushBack adds a new node with the given value to the end of the list.
// It traverses the list to find the last node and sets its next pointer to the new node.
// The time complexity of this operation is O(n), where n is the length of the list.
func (l *SinglyLinkedList[T]) PushBack(x T) {
	node := &Node[T]{
		Data: &x,
		next: nil,
	}
	// if the length is 0, it means that we insert a first node
	if l.length == 0 {
		l.head = node
	} else {
		// oterwise travers the list to find the end
		for n := l.head; n != nil; n = n.next {
			if n.next == nil {
				n.next = node
				break
			}
		}
	}
	l.length++
}
