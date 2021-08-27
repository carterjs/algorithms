package insertionSort

import (
	"fmt"
	"math/rand"
	"testing"
)

// TestExec tests the sort function with slices of random integers (as my assignment dictated)
func TestExec(t *testing.T) {
	type test struct {
		name string
		input []int
		expected []int
	}

	tests := []test{
		{
			name: "basic",
			input: []int{5,3,2,4,6,7,8,9,1},
			expected: []int{1,2,3,4,5,6,7,8,9},
		},
		{
			name: "singleNumber",
			input: []int{0},
			expected: []int{0},
		},
	}

	// Run table tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := Exec(tc.input)

			// compare output to expected
			for i, expected := range tc.expected {
				if output[i] != expected {
					t.Fatalf("Expected %v, got %v", tc.expected, output)
				}
			}
		})
	}
}

// BenchmarkExec runs test cases with random integer permutations of length n=1..10
func BenchmarkExec(b *testing.B) {
	for i := 0; i< b.N; i++ {
		// Generate benchmark for n of 1-10
		for n := 1; n <= 10; n++ {
			b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
				// Sort array with n elements
				Exec(rand.Perm(n))
			})
		}
	}
}