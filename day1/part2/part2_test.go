package part2

import (
	"aoc2020/day1/part1"
	"testing"
)

func TestCalculateFuelNeededForMass(t *testing.T) {
	type args struct {
		module part1.Module
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Mass 12", args: args{module: part1.Module{Mass: 12}}, want: 2},
		{name: "Mass 14", args: args{module: part1.Module{Mass: 14}}, want: 2},
		{name: "Mass 1969", args: args{module: part1.Module{Mass: 1969}}, want: 966},
		{name: "Mass 100756", args: args{module: part1.Module{Mass: 100756}}, want: 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1.CalculateFuelNeededForMass(tt.args.module); got != tt.want {
				t.Errorf("CalculateFuelNeededForMass() = %v, want %v", got, tt.want)
			}
		})
	}
}