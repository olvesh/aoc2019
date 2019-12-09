package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	maxPhases := 5
	maxPhaseVal := 5
	maxPhaseValMod := 5

	idxes := make([]int, 5)
	for i := 0; i < len(idxes); i++ {
		idxes[i] = i
	}

	res := make([]int, 255)

	x := 0
	for i := 0; i < len(res); i = i + 1 {

		var phases []int
		if i == 0 {
			phases = make([]int, maxPhases)
			copy(phases, idxes)
		} /*else {
		    copy()
		  }
		*/

		fmt.Printf("%v: ", i)

		for j := 0; j < maxPhases; j++ {
			phases[j] = (j + x) % maxPhaseValMod

			x++
		}
		x++
		maxPhaseValMod = maxPhaseVal

		fmt.Printf("%v \n", phases)

	}
}

func TestNewAmpRack(t *testing.T) {
	type args struct {
		phases []int
		prog   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "43210", args: args{
			phases: []int{4, 3, 2, 1, 0},
			prog:   []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		}, want: 43210},
		{name: "54321", args: args{
			phases: []int{0, 1, 2, 3, 4},
			prog:   []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		}, want: 54321},
		{name: "65210", args: args{
			phases: []int{1, 0, 4, 3, 2},
			prog:   []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		}, want: 65210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAmpRack(tt.args.phases, tt.args.prog); got != tt.want {
				t.Errorf("NewAmpRack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAmplitude(t *testing.T) {
	type args struct {
		prog []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{name: "65210",
			args: args{prog: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			want: 65210},

		{name: "54321",
			args: args{prog: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			want: 54321},

		{name: "43210",
			args: args{prog: []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			want: 43210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAmplitude(tt.args.prog); got != tt.want {
				t.Errorf("MaxAmplitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
