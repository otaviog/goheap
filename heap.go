// goheap is a simple implementation of the heap data structure and sorting algorithm.
// Its intention is for learning purposes of the Go language.
package goheap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Generic heap structure
type Heap[T constraints.Ordered] struct {
	array     []T
	size      int
	less_func func(lfs, rhs T) bool
}

func (obj Heap[T]) Capacity() int {
	return len(obj.array)
}

func (obj Heap[T]) Len() int {
	return obj.size
}

func New[T constraints.Ordered](capacity int, less_func func(lfs, rhs T) bool) *Heap[T] {
	heap := Heap[T]{make([]T, capacity), 0, less_func}
	return &heap
}

func MakeHeap[T constraints.Ordered](array []T, less_func func(lfs, rhs T) bool) Heap[T] {
	heap := Heap[T]{array, 0, less_func}

	for _, v := range array {
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

func (heap *Heap[T]) Insert(val T) error {
	if heap.size >= heap.Capacity() {
		return errors.New("Heap capacity")
	}
	heap.array[heap.size] = val

	var burble_up func(int)
	burble_up = func(node int) {
		parent_node := parent(node)
		if parent_node >= 0 {
			parent_val := heap.array[parent_node]
			val := heap.array[node]
			if !(heap.less_func(parent_val, val)) {
				heap.array[node], heap.array[parent_node] = heap.array[parent_node], heap.array[node]
				burble_up(parent_node)
			}
		}
	}

	burble_up(heap.size)
	heap.size += 1
	return nil
}

func (heap *Heap[T]) Remove() (removedValue T, err error) {
	if heap.size == 0 {
		var zero T
		return zero, errors.New("Empty heap")
	}

	var burbleDown func(int)
	burbleDown = func(node int) {
		nodeVal := heap.array[node]
		firstChild := firstChildren(node)

		minK := -1
		minVal := nodeVal
		for k := 0; k < 2 && firstChild+k < heap.size; k++ {
			childVal := heap.array[firstChild+k]
			if heap.less_func(childVal, minVal) {
				minK = k
				minVal = childVal
			}
		}

		if minK > -1 {
			child := firstChild + minK
			heap.array[node], heap.array[child] = heap.array[child], nodeVal
			burbleDown(child)
		}
	}

	removedValue = heap.array[0]
	heap.array[0], heap.size = heap.array[heap.size-1], heap.size-1
	burbleDown(0)

	return
}

func HeapSort[T constraints.Ordered](array []T) {
	heap := MakeHeap(array, func(lfs, rhs T) bool {
		return lfs > rhs
	})
	if len(array) < 2 {
		return
	}

	for i := len(array) - 1; i >= 0; i-- {
		array[i], _ = heap.Remove()
	}
}
