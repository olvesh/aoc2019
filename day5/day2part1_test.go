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

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcIntOpsSlice(tt.args.intops); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v: CalcIntOps() = %v, want %v", i, got, tt.want)
			}
		})
	}
}

func TestResultDay2(t *testing.T) {
	var gravityAssistProgramInput = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 9, 23, 1, 5, 23, 27, 1, 27, 9, 31, 1, 6, 31, 35, 2, 35, 9, 39, 1, 39, 6, 43, 2, 9, 43, 47, 1, 47, 6, 51, 2, 51, 9, 55, 1, 5, 55, 59, 2, 59, 6, 63, 1, 9, 63, 67, 1, 67, 10, 71, 1, 71, 13, 75, 2, 13, 75, 79, 1, 6, 79, 83, 2, 9, 83, 87, 1, 87, 6, 91, 2, 10, 91, 95, 2, 13, 95, 99, 1, 9, 99, 103, 1, 5, 103, 107, 2, 9, 107, 111, 1, 111, 5, 115, 1, 115, 5, 119, 1, 10, 119, 123, 1, 13, 123, 127, 1, 2, 127, 131, 1, 131, 13, 0, 99, 2, 14, 0, 0}
	noun := 12
	verb := 2

	res := execIntcode(mapify(gravityAssistProgramInput), noun, verb)

	if !(res == 3409710) {
		t.Errorf("Expected 3409710, got %v", res)
	}

}
