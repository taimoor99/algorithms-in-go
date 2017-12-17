package sort

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

	// split the array in two parts
	left := a[0 : len(a)/2]
	right := a[(len(a) / 2):len(a)]

	left = MergeSort(left)
	right = MergeSort(right)

	aux := make([]int, len(a))
	return Merge(left, right, aux)
}
