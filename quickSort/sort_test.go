package quickSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Basic",
			args: args{
				a: []int{1,4,3},
			},
			want: []int{1,3,4},
		},
		{
			name: "SingleNumber",
			args: args{
				a: []int{1},
			},
			want: []int{1},
		},
		{
			name: "Empty",
			args: args{
				a: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Sort(tt.args.a); !reflect.DeepEqual(tt.args.a, tt.want) {
				t.Errorf("a after Sort() = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}
