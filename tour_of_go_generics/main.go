package main

import (
	"fmt"
	"strings"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Add an element to index position
func (l *List[T]) Insert(i int, x T) *List[T] {
	if i < 0 {
		return l
	}

	if i == 0 {
		return &List[T]{val: x, next: l}
	}

	curr := l
	for counter := 0; counter < i-1; counter++ {
		if curr == nil {
			return l
		}
		curr = curr.next
	}

	if curr == nil {
		return l
	}

	curr.next = &List[T]{val: x, next: curr.next}
	return l
}

// Add to end
func (l *List[T]) Append(x T) *List[T] {
	newNode := &List[T]{val: x}
	if l == nil {
		return newNode
	}

	curr := l
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode

	return l
}

// Removes last element
func (l *List[T]) Pop() *List[T] {
	if l == nil || l.next == nil {
		return nil
	}

	curr := l
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil

	return l
}

func (l *List[T]) String() string {
	if l == nil {
		return "nil"
	}

	var sb strings.Builder
	fmt.Fprint(&sb, "[")
	for curr := l; curr != nil; curr = curr.next {
		fmt.Fprintf(&sb, "%v", curr.val)
		if curr.next != nil {
			fmt.Fprint(&sb, ", ")
		}
	}
	fmt.Fprint(&sb, "]")

	return sb.String()
}

func main() {
	var iList *List[int]

	iList = iList.Append(4)
	iList = iList.Append(5)
	iList = iList.Append(6)
	iList = iList.Insert(0, 3)
	iList = iList.Insert(0, 2)
	iList = iList.Insert(0, 1)
	iList = iList.Pop()

	fmt.Println(iList)
}
