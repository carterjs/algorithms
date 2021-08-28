package insertionSort

import (
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		name string
		input []int
		expected []int
	}

	tests := []test{
		{
			name: "Basic",
			input: []int{5,3,2,4,6,7,8,9,1},
			expected: []int{1,2,3,4,5,6,7,8,9},
		},
		{
			name: "SingleNumber",
			input: []int{0},
			expected: []int{0},
		},
	}

	// Run table tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := Sort(tc.input)

			// compare output to expected
			for i, expected := range tc.expected {
				if output[i] != expected {
					t.Fatalf("Expected %v, got %v", tc.expected, output)
				}
			}
		})
	}
}