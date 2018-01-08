package tests

import (
	"testing"

	"github.com/valentijnnieman/algorithms/sort"
)

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestInsertionSort(t *testing.T) {
	a := []int{10, 5, 9, 2, 4, 3}
	sorted := sort.InsertionSort(a)
	expected := []int{2, 3, 4, 5, 9, 10}

	if !IntArrayEquals(sorted, expected) {
		t.Errorf("A was not sorted correctly with InsertionSort, got: %d, want: %d.", sorted, expected)
	} else {
		t.Log("Insertion sort passed test")
	}
}

func TestMergeSort(t *testing.T) {
	a := []int{10, 5, 9, 2, 4, 3}
	sorted := sort.MergeSort(a)
	expected := []int{2, 3, 4, 5, 9, 10}

	if !IntArrayEquals(sorted, expected) {
		t.Errorf("A was not sorted correctly with MergeSort , got: %d, want: %d.", sorted, expected)
	} else {
		t.Log("Merge sort passed test")
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{10, 5, 9, 2, 4, 3}
	sorted := sort.QuickSort(a)
	expected := []int{2, 3, 4, 5, 9, 10}

	if !IntArrayEquals(sorted, expected) {
		t.Errorf("A was not sorted correctly with MergeSort , got: %d, want: %d.", sorted, expected)
	} else {
		t.Log("Merge sort passed test")
	}
}
