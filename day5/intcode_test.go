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
		want []int
	}{
		{
			name: "1101,100,-1,4,0",
			args: args{intops: []int{1101, 100, -1, 4, 0}},
			want: []int{1101, 100, -1, 4, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcIntOpsSlice(tt.args.intops); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcIntOpsSlice() = %v, want %v", got, tt.want)
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
				opcode:     instructions[1],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInstruction(tt.args.inst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
