package part1

import "testing"

func TestCalculateFuelNeededForMassPart1(t *testing.T) {
	type args struct {
		module Module
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Mass 12", args: args{module: Module{Mass: 12}}, want: 2},
		{name: "Mass 14", args: args{module: Module{Mass: 14}}, want: 2},
		{name: "Mass 1969", args: args{module: Module{Mass: 1969}}, want: 654},
		{name: "Mass 100756", args: args{module: Module{Mass: 100756}}, want: 33583},
		{name: "Mass 100756", args: args{module: Module{Mass: 100756}}, want: 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFuelNeededForMass(tt.args.module); got != tt.want {
				t.Errorf("CalculateFuelNeededForMass() = %v, want %v", got, tt.want)
			}
		})
	}
}