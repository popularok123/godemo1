package main

import (
	"fmt"
	"iter"
	"slices"
)

type GenericList struct{}

func SlieceIndex[S ~[]E, E comparable](s S, e E) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Append(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yieldddd func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yieldddd(e.val) {
				return
			}
		}
	}
}

func getFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func (gl *GenericList) Do() {

	var s = []string{"one", "two", "three"}

	fmt.Println("index of zoo:", SlieceIndex(s, "two"))

	_ = SlieceIndex[[]string, string](s, "three")

	lst := List[int]{}

	lst.Append(1)
	lst.Append(13)
	lst.Append(123)
	fmt.Println("list:", lst.AllElements())

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range getFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
