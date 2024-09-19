package slices

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	// Declaration and initialization of a slice
	// The size of a slice can grow and shrink depending on your use case
	names := make([]string, 4, 10)

	// Assigning values to slice elements
	names[0] = "Samuel"
	names[1] = "David"
	names[2] = "Stanley"
	names[3] = "James"

	// Accessing slice elements via indexing
	fmt.Println("Element at index 3:", names[3])

	// Iterating over slice elements
	for i := 0; i < len(names); i++ {
		fmt.Printf("Element at index %d: %s\n", i, names[i])
	}

	// Declaration and initialization of an slice with values
	ages := []int64{20, 12, 34, 19}

	// Length of the slice
	fmt.Println("Length of ages slice:", len(ages))

	// Capacity of the slice
	fmt.Println("Length of ages slice:", cap(ages))

	// Modifying slices
	newAge := []int64{46, 32, 13, 11}

	// Length of the slice before adding new ages
	fmt.Println("Length of ages slice before adding new ages:", len(ages))

	ages = append(ages, newAge...)
	// Length of the slice after adding new ages
	fmt.Println("Length of ages slice after adding new ages:", len(ages))

	// Value of ages slice after appending newAge
	fmt.Println("Value of ages slice after adding new ages:", ages)

}
