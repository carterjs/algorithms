package mergeSort

// Merge combines two sorted integer slices
func Merge(a []int, b []int) []int {
	var ai, bi int
	r := make([]int, len(a) + len(b))

	// loop enough times to operate on every element
	for i := range r {
		// default to "a" if "b" is done
		if bi > len(b) - 1 {
			r[i] = a[ai]
			ai++
			continue
		}

		// default to "b" if "a" is done
		if ai > len(a) - 1 {
			r[i] = b[bi]
			bi++
			continue
		}

		// Compare
		if a[ai] <= b[bi] {
			// "a[ai]" is smaller, or they're the same
			r[i] = a[ai]
			ai++
		} else {
			// "b[bi]" is smaller
			r[i] = b[bi]
			bi++
		}
	}

	return r
}