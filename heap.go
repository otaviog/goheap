// goheap is a simple implementation of the heap data structure and sorting algorithm.
// Its intention is for learning purposes of the Go language.
package goheap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Generic heap structure
type Heap[T any] struct {
	slice    []T
	size     int
	lessFunc func(lfs, rhs T) bool
}

// The current capacity to store elements.
func (obj Heap[T]) Capacity() int {
	return cap(obj.slice)
}

// The number of elements inside the heap.
func (obj Heap[T]) Len() int {
	return obj.size
}

// Creates a new heap data structure ordered by the lesser function lessFunc and
// with optional starting values.
func New[T any](lessFunc func(lfs, rhs T) bool, values ...T) *Heap[T] {
	heap := Heap[T]{make([]T, 0), 0, lessFunc}
	for _, v := range values {
		heap.Insert(v)
	}
	return &heap
}

// Makes an existing slice an Heap ordered according to the lessFunc function.
func MakeHeap[T any](slice []T, lessFunc func(lfs, rhs T) bool) Heap[T] {
	heap := Heap[T]{slice, 0, lessFunc}

	for _, v := range slice {
		heap.Insert(v)
	}

	return heap
}

func parent(node int) int {
	if node == 0 {
		return -1
	}
	return (node+1)/2 - 1
}

func firstChildren(node int) int {
	return node*2 + 1
}

// Inserts a value into the heap. If its internal slice is full then it will append the element
// which breaks the pointer to the old slice.
// Its time complexity is O(log N) swaps where N is the heap size.
func (heap *Heap[T]) Insert(value T) {
	if heap.size >= len(heap.slice) {
		heap.slice = append(heap.slice, value)
	} else {
		heap.slice[heap.size] = value
	}

	var burbleUp func(int)
	burbleUp = func(node int) {
		parentNode := parent(node)
		if parentNode >= 0 {
			parentValue := heap.slice[parentNode]
			val := heap.slice[node]
			if !(heap.lessFunc(parentValue, val)) {
				heap.slice[node], heap.slice[parentNode] = heap.slice[parentNode], heap.slice[node]
				burbleUp(parentNode)
			}
		}
	}

	burbleUp(heap.size)
	heap.size += 1
}

// Remove the lesser value from the heap, example:
//  heap := MakeHeap([]int{8, 9, 4, 2, 7}, func(a, b int) bool { return a < b })
//  value, _ := heap.Remove()
//  value, _ := heap.Remove()
// It should output 2 and 4. If the heap is empty, then returns an error.
// Its time complexity is O(log N) where N is the heap size.
func (heap *Heap[T]) Remove() (removedValue T, err error) {
	if heap.size == 0 {
		var zero T
		return zero, errors.New("Empty heap")
	}

	var burbleDown func(int)
	burbleDown = func(node int) {
		nodeVal := heap.slice[node]
		firstChild := firstChildren(node)

		minK := -1
		minVal := nodeVal
		for k := 0; k < 2 && firstChild+k < heap.size; k++ {
			childVal := heap.slice[firstChild+k]
			if heap.lessFunc(childVal, minVal) {
				minK = k
				minVal = childVal
			}
		}

		if minK > -1 {
			child := firstChild + minK
			heap.slice[node], heap.slice[child] = heap.slice[child], nodeVal
			burbleDown(child)
		}
	}

	removedValue = heap.slice[0]
	heap.slice[0], heap.size = heap.slice[heap.size-1], heap.size-1
	burbleDown(0)

	return
}

// Sorts a slice in increasing order. Its time complexity is O(N log N) where N is the size of the
// slice.
func HeapSort[T constraints.Ordered](slice []T) {
	heap := MakeHeap(slice, func(lfs, rhs T) bool {
		return lfs > rhs
	})
	if len(slice) < 2 {
		return
	}

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i], _ = heap.Remove()
	}
}
