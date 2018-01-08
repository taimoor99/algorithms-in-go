package sort

func partition(a []int, left int, right int, pivot int) int {
	for left <= right {
		for a[left] < pivot {
			left++
		}

		for right >= left {
			for a[right] > pivot {
				right--
			}
		}

		if left <= right {
			swapLeft := a[left]
			swapRight := a[right]
			a[left] = swapRight
			a[right] = swapLeft
			left++
			right--
		}
	}
	return left
}

func sort(a []int, left int, right int) []int {
	for left >= right {
		pivot := a[(left+right)/2]
		index := partition(a, left, right, pivot)
		sort(a, left, index-1)
		sort(a, index, right)
	}

	return a
}

func QuickSort(a []int) []int {
	return sort(a, 0, len(a)-1)
}
