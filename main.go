package main

import (
	"fmt"
	"github.com/valentijnnieman/algorithms/sort"
)

func main() {
	a := []int{10, 5, 9, 2, 4, 3}
	fmt.Println("a: ", a)
	fmt.Println("InsertionSort(a): ", sort.InsertionSort(a))
	a = []int{10, 5, 9, 2, 4, 3, 91, 44, 20, 12, 16, 4}
	fmt.Println("a: ", a)
	fmt.Println("MergeSort(a): ", sort.MergeSort(a))
}
