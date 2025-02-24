package main

import (
	"fmt"
	"reflect"
	"slices"

	bubblesort "github.com/jimtrung/learn-dsa/go/sortalgorithms/BubbleSort"
	insertionsort "github.com/jimtrung/learn-dsa/go/sortalgorithms/InsertionSort"
	selectionsort "github.com/jimtrung/learn-dsa/go/sortalgorithms/SelectionSort"
)

func TestSort(f func([]int)) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}}, // Normal case
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // Already sorted
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, // Reverse sorted
		{[]int{3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3}}, // Duplicates
		{[]int{1}, []int{1}},                         // Single element
		{[]int{}, []int{}},                           // Empty list
	}

	for i, test := range tests {
		arr := slices.Clone(test.input)
        f(arr)

		if !reflect.DeepEqual(arr, test.expected) {
			fmt.Printf(
				"Test %d failed: got %v, expected %v\n",
				i + 1, arr, test.expected,
			)
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}

func main() {
    fmt.Printf("Testing Bubble Sort\n")
	TestSort(bubblesort.BubbleSort)
    fmt.Println()

    fmt.Printf("Testing Selection Sort\n")
    TestSort(selectionsort.SelectionSort)
    fmt.Println()

    fmt.Printf("Testing Insertion Sort\n")
    TestSort(insertionsort.InsertionSort)
    fmt.Println()
}
