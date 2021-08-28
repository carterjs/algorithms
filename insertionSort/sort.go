package insertionSort

// Sort runs the insertion sort algorithm on the input array and returns a sorted array
func Sort(a []int) []int {
	// Make a copy of the slice
	r := make([]int, len(a))
	copy(r, a)

	for i := 1; i < len(r); i++ {
		// selected element
		k := r[i]

		// make room for k in the sorted section
		j := i - 1
		for j >= 0 && r[j] > k {
			// Move this element up
			r[j+1] = r[j]
			j -= 1
		}

		// place k at the proper location
		r[j+1] = k
	}

	return r
}
