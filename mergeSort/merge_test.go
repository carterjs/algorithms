package mergeSort

import "testing"

func TestMerge(t *testing.T) {
	type test struct {
		name string
		input1 []int
		input2 []int
		expected []int
	}

	tests := []test{
		{
			name: "Basic",
			input1: []int{1,3,5,7,9},
			input2: []int{2,4,6,8},
			expected: []int{1,2,3,4,5,6,7,8,9},
		},
		{
			name: "OneEmpty",
			input1: []int{},
			input2: []int{1,2,3},
			expected: []int{1,2,3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := Merge(tc.input1, tc.input2)

			for i, expected := range tc.expected {
				if output[i] != expected {
					t.Fatalf("Expected %v, got %v", tc.expected, output)
				}
			}
		})
	}
}