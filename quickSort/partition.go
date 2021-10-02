package quickSort

// Partition chooses a pivot and puts all other elements in the correct position relative to the pivot
func Partition(a []int, p int, r int) int {
	// i is the last element <= pivot
	i := p - 1

	// Loop through all elements
	for j := p; j < r; j++ {
		if a[j] <= a[r] {
			// Element is <= pivot, increase i and put a[j] at index i
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	// Put pivot in the middle
	a[i + 1], a[r] = a[r], a[i + 1]

	// Return pivot index
	return i + 1
}
