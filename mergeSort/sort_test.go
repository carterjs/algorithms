package mergeSort

import (
	"testing"
)

func TestSort(t *testing.T) {
	type test struct {
		name string
		input []int
		expected []int
	}

	tests := []test{
		{
			name: "Basic",
			input: []int{1,4,3,6,7,4,6,8,3},
			expected: []int{1,3,3,4,4,6,6,7,8},
		},
		{
			name: "SingleNumber",
			input: []int{6},
			expected: []int{6},
		},
		{
			name: "Empty",
			input: []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := Sort(tc.input)

			for i, expected := range tc.expected {
				if output[i] != expected {
					t.Fatalf("Expected %v, got %v", tc.expected, output)
				}
			}
		})
	}
}