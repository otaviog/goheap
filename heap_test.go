package goheap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestParent(t *testing.T) {
	if parent(0) != -1 {
		t.Errorf("Got %d, expected -1", parent(0))
	}
}

func isHeap[T constraints.Ordered](heap Heap[T]) bool {

	var _rec func(int) bool
	_rec = func(node int) bool {
		children := firstChildren(node)

		for k := 0; k < 2; k++ {
			child := children + k
			if child < heap.Len() {
				if heap.array[child] < heap.array[node] {
					return false
				}
			}
		}

		for k := 0; k < 2; k++ {
			child := children + k
			if child < heap.size {
				if !_rec(child) {
					return false
				}
			}
		}
		return true
	}

	return _rec(0)
}

func TestInsertion(t *testing.T) {
	assert := assert.New(t)
	array := []int{8, 2, 3, 9, 5, 4}
	heap := Heap[int]{array, 0, func(lfs, rhs int) bool { return lfs < rhs }}

	assert.Equal(0, heap.Len())
	assert.Equal(len(array), heap.Capacity())
	for i, v := range array {
		err := heap.Insert(v)

		assert.Nil(err)
		assert.Equal(i+1, heap.Len())
		assert.True(isHeap(heap))
	}
	assert.NotNil(heap.Insert(5))
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	heap := MakeHeap([]int{8, 2, 1, 4, 5, 4}, func(lfs, rhs int) bool { return lfs < rhs })

	removeOrder := make([]int, heap.Len())
	for i := 0; i < len(removeOrder); i++ {
		removeOrder[i], _ = heap.Remove()
	}

	assert.True(sort.IntsAreSorted(removeOrder), "Bug in remove, got", removeOrder)

	_, err := heap.Remove()
	assert.NotNil(err)
}

func TestHeapsort(t *testing.T) {
	assert := assert.New(t)

	array := []int{8, 2, 3, 9, 5, 4}
	HeapSort(array)
	assert.True(sort.IntsAreSorted(array), "bug in heapsort")

	array = []int{}
	HeapSort(array)

	array = []int{4}
	HeapSort(array)
}

func BenchmarkHeapsort(b *testing.B) {
	rand := rand.New(rand.NewSource(4444))
	array := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(array); i++ {
			array[i] = rand.Int()
		}
		HeapSort(array)
	}
}

func BenchmarkQuicksort(b *testing.B) {
	rand := rand.New(rand.NewSource(4444))
	array := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(array); i++ {
			array[i] = rand.Int()
		}
		sort.IntsAreSorted(array)
	}
}
