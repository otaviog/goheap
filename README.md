# Go Heap 

This is a simple implementation of a heap data structure using Go and its generics.
Not intended for production.

## Usage

Install:

```shell
go get github.com/otaviog/goheap
```

Create a new heap:

```go
heap := goheap.New(func(a, b int) bool { return a < b }, 4, 1, 2, 4, 8, 0, 3)
heap.Insert(5)
lowest, _ = heap.Remove()
// Outputs 0
```

Make an heap from an already existing slice:

```go
slice := []int{8, 2, 1, 4, 5, 4}
heap := goheap.MakeHeap(slice, func(lfs, rhs int) bool { return lfs < rhs })
lowest := heap.Remove()
// Outputs 1
heap.Insert(10) // Keeps using the initial slice array, slice[5] = 10
heap.Insert(11) // Overflows the current slice, a new one is append
```

Sort orderable types (`constraints.Ordered`):

```go
unorderedSlice := [int]...
HeapSort(unorderedSlice)
```

Usage with custom types:

```go
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
		name:  "Chris",
		phone: "555-2121",
	},
	{
		age:   42,
		name:  "Rochelle",
		phone: "555-4421",
	},
}
heap := MakeHeap(persons, func(p1, p2 Person) bool {
	return p1.age < p2.age
})
heap.Insert(Person{age: 13, name: "Vicent", phone: "555-4211"})
person, _ := heap.Remove()
```

[Reference API](https://otaviog.github.io/goheap/).

# Benchmarks

Some benchmarks for the sake of completeness.

## Integer sorting

For sorting integers, GoHeap's `Heapsort` with recursion had poor results when compared with `sort.Ints` from Go's standard library:

Next are the results for the version v1.0 with recursion:

```shell
go test -bench=. -count=5
goos: linux
goarch: amd64
pkg: github.com/otaviog/goheap
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkHeapsort                  16          73611175 ns/op
BenchmarkHeapsort                  15          73011714 ns/op
BenchmarkHeapsort                  15          73200773 ns/op
BenchmarkHeapsort                  16          75392815 ns/op
BenchmarkHeapsort                  15          72352460 ns/op
BenchmarkStdSort                  330           3618156 ns/op
BenchmarkStdSort                  336           3655944 ns/op
BenchmarkStdSort                  334           3624276 ns/op
BenchmarkStdSort                  328           3711837 ns/op
BenchmarkStdSort                  334           3632133 ns/op
```

And next are the results for the version v1.1 without recursion, still far from the standard library one:

```shell
go test -bench=. -count=5
goos: linux
goarch: amd64
pkg: github.com/otaviog/goheap
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkHeapsort                  26          46133913 ns/op
BenchmarkHeapsort                  27          45261210 ns/op
BenchmarkHeapsort                  26          44689030 ns/op
BenchmarkHeapsort                  25          46688725 ns/op
BenchmarkHeapsort                  26          47114640 ns/op
BenchmarkStdSort                  333           3629980 ns/op
BenchmarkStdSort                  316           3568491 ns/op
BenchmarkStdSort                  340           3523823 ns/op
BenchmarkStdSort                  336           3540776 ns/op
BenchmarkStdSort                  340           3569386 ns/op
```
