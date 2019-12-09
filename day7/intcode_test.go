package main

import (
	"reflect"
	"testing"
)

func TestCalcIntOpsSlice(t *testing.T) {
	type args struct {
		intops []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "1101,100,-1,4,0",
			args: args{intops: []int{1101, 100, -1, 4, 0}},
			want: map[int]int{0: 1101, 1: 100, 2: -1, 3: 4, 4: 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computer := NewIntcodeComputer(tt.args.intops)
			if got := computer.Exec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntcodeComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Fails, but is ok :-/
func TestNewInstruction(t *testing.T) {
	type args struct {
		inst int
	}
	tests := []struct {
		name string
		args args
		want Instruction
	}{
		{
			name: "1101",
			args: args{inst: 1101},
			want: Instruction{
				paramModes: []int{1, 1, 0, 0},
				opcode:     ops[1],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInstruction(tt.args.inst, ops); &got.opcode != &tt.want.opcode || !reflect.DeepEqual(got.paramModes, tt.want.paramModes) {
				t.Errorf("newInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
