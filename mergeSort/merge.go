package mergeSort

// Merge combines two sorted integer slices
func Merge(a []int, b []int) []int {
	var ai, bi int
	var r []int

	// loop enough times to operate on every element
	for {
		// break out if we've gone through them all
		if ai > len(a) - 1 && bi > len(b) - 1 {
			break
		}

		// default to "a" if "b" is done
		if bi > len(b) - 1 {
			r = append(r, a[ai])
			ai++
			continue
		}

		// default to "b" if "a" is done
		if ai > len(a) - 1 {
			r = append(r, b[bi])
			bi++
			continue
		}

		// Compare
		if a[ai] <= b[bi] {
			// "a[ai]" is smaller, or they're the same
			r = append(r, a[ai])
			ai++
		} else {
			// "b[bi]" is smaller
			r = append(r, b[bi])
			bi++
		}
	}

	return r
}