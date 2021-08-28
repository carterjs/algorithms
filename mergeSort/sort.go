package mergeSort

func Sort(a []int) []int {
	// Stop condition - only one element
	if len(a) <= 1 {
		return a
	}

	// Sort left side
	l := Sort(a[:len(a)/2])

	// Sort right side
	r := Sort(a[len(a)/2:])

	// Merge sides
	return Merge(l, r)
}
