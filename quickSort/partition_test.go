package quickSort

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	type args struct {
		a []int
		p int
		r int
	}
	type want struct {
		a []int
		pivot int
	}
	tests := []struct {
		name string
		args   args
		result want
	}{
		{
			name: "Basic",
			args: args{
				a: []int{1,5,3,2,4},
				p: 0,
				r: 4,
			},
			result: want{
				a: []int{1,3,2,5,4},
				pivot: 3,
			},
		},
		{
			name: "Sorted",
			args: args{
				a: []int{1,2,3,4,5},
				r: 0,
				p: 4,
			},
			result: want{
				a: []int{1,2,3,4,5},
				pivot: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pivot := Partition(tt.args.a, tt.args.p, tt.args.r)

			// Check pivot
			if pivot != tt.result.pivot {
				t.Errorf("Partition() = %v, want %v", pivot, tt.result.pivot)
			}

			// Check end want
			if reflect.DeepEqual(tt.args.a, tt.result.a) {
				t.Errorf("a after Partition() = %v, want %v", tt.args.a, tt.result.a)
			}
		})
	}
}
