package main

import "testing"

func TestPassDoubleDigit(t *testing.T) {
	type args struct {
		pass []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "111111", args: args{pass: []byte("111111")}, want: false},
		{name: "111122", args: args{pass: []byte("111122")}, want: true},
		{name: "112233", args: args{pass: []byte("112233")}, want: true},
		{name: "123444", args: args{pass: []byte("123444")}, want: false},
		{name: "223450", args: args{pass: []byte("223450")}, want: true},
		{name: "123789", args: args{pass: []byte("123789")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CriteriaDoubleDigit(tt.args.pass); got != tt.want {
				t.Errorf("CriteriaDoubleDigit() = %v, want %v", got, tt.want)
			}

		})
	}
}
func TestPassSorted(t *testing.T) {
	type args struct {
		pass []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "111111", args: args{pass: []byte("111111")}, want: true},
		{name: "112233", args: args{pass: []byte("112233")}, want: true},
		{name: "223450", args: args{pass: []byte("223450")}, want: false},
		{name: "123789", args: args{pass: []byte("123789")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CriteriaSorted(tt.args.pass); got != tt.want {
				t.Errorf("CriteriaSorted() = %v, want %v", got, tt.want)
			}

		})
	}
}
