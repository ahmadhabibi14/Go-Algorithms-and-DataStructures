package main

import "fmt"

func reverseArray(arr []int) []int {
	var n = len(arr) - 1
	for i := 0; i < n; i, n = i+1, n-1 {
		arr[i], arr[n] = arr[n], arr[i]
	}
	return arr
}
func main() {
	var a = []int{1, 4, 3, 2}

	fmt.Println("Original Array = ", a)
	fmt.Println("Reversed Array = ", reverseArray(a))
}
