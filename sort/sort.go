package sort

import "fmt"

func InsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i -= 1
		}
		a[i+1] = key
	}
	return a
}

func Merge(left []int, right []int, aux []int) []int {
	i := 0
	j := 0
	for k := 0; k < len(aux); k++ {
		if i >= len(left) {
			// reached end of left
			aux[k] = right[j]
			j++
		} else if j >= len(right) {
			// reached end of right
			aux[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			aux[k] = left[i]
			i++
		} else {
			aux[k] = right[j]
			j++
		}
	}
	return aux
}

func MergeSort(a []int) []int {
	// let's test out merge
	if len(a) <= 1 {
		return a
	}

	fmt.Printf("a: %d \n", a)

	// split the array in two parts
	left := a[0 : len(a)/2]
	right := a[(len(a) / 2):len(a)]

	left = MergeSort(left)
	right = MergeSort(right)

	aux := make([]int, len(a))
	return Merge(left, right, aux)
}

/*func QuickSort(a []int) []int {*/
//// STOP if no elements
//if len(a) <= 0 {
//return
//}

//// first shuffle the array
//for i := range a {
//j := rand.Intn(i + 1)
//a[i], a[j] = a[j], a[i]
//}

//// then start partitioning
//pivot := 0
//for i := 1; i < len(a); i++ {
//for j := len(a); j > i; j-- {
//if a[i] > a[pivot] {

//}
//if a[j] < a[pivot] {

//}
//}
//}
/*}*/
