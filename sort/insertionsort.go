package sort

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
