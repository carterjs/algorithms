package quickSort

// Sort sorts an array in place using quicksort
func Sort(a []int) {
	// If the subarray isn't empty
	if len(a) > 1 {
		// Partition - get pivot q
		q := Partition(a, 0, len(a) - 1)

		// Call recursively on all elements <= q
		Sort(a[:q])

		// Call recursively on all elements > q
		Sort(a[q:])
	}
}