package sorting

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var arr []int = []int{12, 11, 13, 5, 6, 4, 7, 10, 2, 3, 1, 9, 8}
	var n int = len(arr)

	insertionSort(arr, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr []int = []int{12, 11, 13, 5, 6, 4, 7, 10, 2, 3, 1, 9, 8}
		var n int = len(arr)

		insertionSort(arr, n)
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println()
	}
}

func insertionSort(arr []int, n int) {
	var i, key, j int
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j = j -1
		}
		arr[j + 1] = key
	}
}