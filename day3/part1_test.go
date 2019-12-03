package main

import (
	"image"
	"reflect"
	"testing"
)

// TODO What about parallel lines?
func TestSegment_not_intersects(t *testing.T) {
	segment1 := NewSegment("R10", image.Point{})
	segment2 := NewSegment("D10", image.Point{X: 11, Y: 5})

	intersects, p := segment1.intersects(segment2)
	if intersects {
		t.Logf("%+v did not intersect %+v at %+v", segment1, segment2, p)
		t.Fail()
	}

	intersects, p = segment2.intersects(segment1)
	if intersects {
		t.Logf("%+v did not intersect %+v at %+v", segment1, segment2, p)
		t.Fail()
	}

}

func TestSegment_intersects(t *testing.T) {
	segment1 := NewSegment("R10", image.Point{})
	segment2 := NewSegment("D10", image.Point{X: 5, Y: 5})

	intersects, _ := segment1.intersects(segment2)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

	intersects, _ = segment2.intersects(segment1)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

}

func TestSegment_intersectionPointXaxis(t *testing.T) {
	segment1 := NewSegment("R10", image.Point{0, 0})
	segment2 := NewSegment("D10", image.Point{X: 5, Y: 5})

	expected := image.Point{X: 5}

	intersects, point := segment1.intersects(segment2)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

	if expected != point {
		t.Logf("%+v != %+v", point, expected)
		t.Fail()
	}

	intersects, point = segment2.intersects(segment1)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

	if expected != point {
		t.Logf("%+v != %+v", point, expected)
		t.Fail()
	}

}

func TestSegment_intersectionPointYaxis(t *testing.T) {
	segment1 := NewSegment("U10", image.Point{0, 0})
	segment2 := NewSegment("L10", image.Point{X: 5, Y: 5})

	expected := image.Point{Y: 5}

	intersects, point := segment1.intersects(segment2)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

	if expected != point {
		t.Logf("%+v != %+v", point, expected)
		t.Fail()
	}

	intersects, point = segment2.intersects(segment1)
	if !intersects {
		t.Logf("%+v did not intersect %+v", segment1, segment2)
		t.Fail()
	}

	if expected != point {
		t.Logf("%+v != %+v", point, expected)
		t.Fail()
	}

}

func TestDistance(t *testing.T) {
	type args struct {
		w1 Wire
		w2 Wire
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "distane empty", args: args{
			w1: Wire{},
			w2: Wire{},
		}, want: -1},
		{name: "distance 159", args: args{
			w1: NewWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			w2: NewWire("U62,R66,U55,R34,D71,R55,D58,R83"),
		}, want: 159},

		{name: "distance 135", args: args{
			w1: NewWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			w2: NewWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
		}, want: 135},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.w1, tt.args.w2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	type args struct {
		w1 Wire
		w2 Wire
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "steps empty", args: args{
			w1: Wire{},
			w2: Wire{},
		}, want: -1},
		{name: "steps 30", args: args{
			w1: NewWire("R8,U5,L5,D3"),
			w2: NewWire("U7,R6,D4,L4"),
		}, want: 30},
		{name: "steps 610", args: args{
			w1: NewWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			w2: NewWire("U62,R66,U55,R34,D71,R55,D58,R83"),
		}, want: 610},

		{name: "steps 410", args: args{
			w1: NewWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			w2: NewWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
		}, want: 410},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Steps(tt.args.w1, tt.args.w2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegment_intersects1(t *testing.T) {

	type args struct {
		seg1 Segment
		seg2 Segment
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 image.Point
	}{
		{
			name: "(0,0)/(0,75) vs (66,62)/(66,117)",
			args: args{
				seg1: NewSegment("U55", image.Point{X: 66, Y: 62}),
				seg2: NewSegment("R75", image.Point{}),
			},
			want:  false,
			want1: image.Point{},
		},
		{
			name: "(0,0)/(0,10) vs (5,5)/(5,-5)",
			args: args{
				seg1: NewSegment("D10", image.Point{X: 5, Y: 5}),
				seg2: NewSegment("R10", image.Point{}),
			},
			want:  true,
			want1: image.Point{5, 0},
		},

		{
			name: "(5,5)/(5,-5) vs (0,0)/(0,10) ",
			args: args{
				seg1: NewSegment("R10", image.Point{}),
				seg2: NewSegment("D10", image.Point{X: 5, Y: 5}),
			},
			want:  true,
			want1: image.Point{5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, got1 := tt.args.seg1.intersects(tt.args.seg2)
			if got != tt.want {
				t.Errorf("intersects() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("intersects() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
