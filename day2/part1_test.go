package main

import (
	"reflect"
	"testing"
)

func TestCalcIntOps(t *testing.T) {
	type args struct {
		intops []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1,9,10,3,2,3,11,0,99,30,40,50",
			args: args{intops: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}},
			want: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{name: "1,0,0,0,99",
			args: args{intops: []int{1, 0, 0, 0, 99}},
			want: []int{2, 0, 0, 0, 99}},
		{name: "2,4,4,5,99,0",
			args: args{intops: []int{2, 4, 4, 5, 99, 0}},
			want: []int{2, 4, 4, 5, 99, 9801}},
		{name: "1,1,1,4,99,5,6,0,99",
			args: args{intops: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
			want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcIntOpsSlice(tt.args.intops); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcIntOps() = %v, want %v", got, tt.want)
			}
		})
	}
}
