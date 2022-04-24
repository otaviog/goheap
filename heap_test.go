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
				if heap.slice[child] < heap.slice[node] {
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

func TestMakeHeap(t *testing.T) {
	assert := assert.New(t)
	heap := MakeHeap([]int{8, 2, 3, 9, 5, 4}, func(a, b int) bool { return a < b })
	assert.True(isHeap(heap))
	assert.Equal(6, heap.Len())
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	heap := New(func(a, b int) bool { return a < b }, 4, 1, 2, 4, 8, 0, 3)
	assert.True(isHeap(*heap))
	assert.Equal(7, heap.Len())
}
func TestInsertion(t *testing.T) {
	assert := assert.New(t)
	array := []int{8, 2, 3, 9, 5, 4}
	heap := Heap[int]{array, 0, func(lfs, rhs int) bool { return lfs < rhs }}

	assert.Equal(0, heap.Len())
	assert.Equal(len(array), heap.Capacity())
	for i, v := range append(array, []int{4, 9, 10, 11}...) {
		heap.Insert(v)

		assert.Equal(i+1, heap.Len())
		assert.True(isHeap(heap))
	}

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
	assert.True(sort.IntsAreSorted(array), "Bug in heapsort")

	array = []int{}
	HeapSort(array)

	array = []int{4}
	HeapSort(array)

	for _, seed := range []int64{5212, 123123, 129852, 982811} {
		rand := rand.New(rand.NewSource(seed))
		unorderedSlice := make([]int, 1000000)
		for i := 0; i < len(unorderedSlice); i++ {
			unorderedSlice[i] = rand.Int()
		}
		HeapSort(unorderedSlice)
		assert.True(sort.IntsAreSorted(unorderedSlice), "Bug in heapsort")
	}

	increasingOrder := [10000]int{}
	for i := range increasingOrder {
		increasingOrder[i] = i
	}

	HeapSort(increasingOrder[:])
	assert.True(sort.IntsAreSorted(increasingOrder[:]), "Bug in heapsort [increasing sorted slice")

	decreasingOrder := [10000]int{}
	for i := range decreasingOrder {
		decreasingOrder[len(decreasingOrder)-i-1] = i
	}
	HeapSort(decreasingOrder[:])
	assert.True(sort.IntsAreSorted(decreasingOrder[:]), "Bug in heapsort [decreasing sorted slice")
}
func TestCustomType(t *testing.T) {
	assert := assert.New(t)

	type Person struct {
		age   int
		name  string
		phone string
	}
	var persons = []Person{
		{
			age:   45,
			name:  "Julius",
			phone: "555-5755",
		},
		{
			age:   12,
			name:  "Cris",
			phone: "555-2121",
		},
		{
			age:   42,
			name:  "Rochele",
			phone: "555-4421",
		},
	}

	heap := MakeHeap(persons, func(p1, p2 Person) bool {
		return p1.age < p2.age
	})
	heap.Insert(Person{age: 13, name: "Vicent", phone: "555-4211"})

	person, _ := heap.Remove()
	assert.Equal(12, person.age)

	person, _ = heap.Remove()
	assert.Equal(13, person.age)
}

func benchmarkSort(b *testing.B, sort_func func([]int)) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(4444))
	unorderedSlice := make([]int, 1000000)
	for i := 0; i < len(unorderedSlice); i++ {
		unorderedSlice[i] = rand.Int()
	}

	temporary_slice := make([]int, len(unorderedSlice))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(unorderedSlice, temporary_slice)
		b.StartTimer()

		sort_func(temporary_slice)
	}
}

func BenchmarkHeapsort(b *testing.B) {
	benchmarkSort(b, HeapSort[int])
}

func BenchmarkStdSort(b *testing.B) {
	benchmarkSort(b, sort.Ints)
}
