package sorting

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	var heap = NewHeap()

	fmt.Println("Unsorted\t:", array)

	heap.HeapSort(array)

	fmt.Println("Sorted\t\t:", array)
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
		var heap = NewHeap()

		heap.HeapSort(array)
		fmt.Println("Sorted with HeapSort \t:", array)
	}
}

type Heap struct {}

func NewHeap() *Heap { return &Heap{} }

func (h *Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (h *Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}

func (h *Heap) Heapify(array []int, root, length int) {
	var max int = root
	var l, r = h.Left(array, root), h.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		h.Heapify(array, max, length)
	}
}

func (h *Heap) RemoveTop(array []int, length int) {
	var lastIndex int = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]

	h.Heapify(array, 0, lastIndex)
}

func (h *Heap) BuildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		h.Heapify(array, i, len(array))
	}
}

func (h *Heap) HeapSort(array []int) {
	h.BuildHeap(array)

	for length := len(array); length > 1; length-- {
		h.RemoveTop(array, length)
	}
}