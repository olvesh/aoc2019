package main

import (
	"image"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Segment struct {
	Direction    string
	Polarisation string
	Length       int
	Start        image.Point
	Stop         image.Point
	rect         image.Rectangle
}

const hor = "RL"
const ver = "UD"

func NewSegment(segment string, start image.Point) Segment {
	length, err := strconv.Atoi(segment[1:])
	if err != nil {
		log.Fatal(err)
	}
	s := Segment{
		Direction: string(segment[0]),
		Length:    length,
		Start:     start,
	}
	s.Stop = s.RelativeTo(s.Start)
	if strings.Contains(hor, s.Direction) {
		s.Polarisation = hor
	} else {
		s.Polarisation = ver
	}
	s.rect = image.Rectangle{
		Min: start,
		Max: s.Stop,
	}

	return s
}

func (s Segment) SetStop(relativeTo image.Point) {
	s.Stop = s.RelativeTo(relativeTo)
}

func (s Segment) RelativeTo(point image.Point) image.Point {
	if s.Direction == "R" {
		point.X = point.X + s.Length
	} else if s.Direction == "L" {
		point.X = point.X - s.Length
	} else if s.Direction == "U" {
		point.Y = point.Y + s.Length
	} else if s.Direction == "D" {
		point.Y = point.Y - s.Length
	}
	return point
}

func (s Segment) intersects(seg2 Segment) (bool, image.Point) {
	if opositeDirections(s, seg2) {

		if b, point := pointsIntersect(s, seg2); b {
			return b, point
		}
		if b, point := pointsIntersect(seg2, s); b {
			return b, point
		}
	}
	return false, image.Point{}
}

func pointsIntersect(seg1 Segment, seg2 Segment) (bool, image.Point) {
	if seg1.Start.X < seg2.Stop.X && seg2.Start.X < seg1.Stop.X ||
		seg1.Start.Y < seg2.Stop.Y && seg2.Start.Y < seg1.Stop.Y {
		return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg2.Start.Y}
	}

	//if seg1.Start.X > seg2.Start.X && seg1.Start.X < seg2.Stop.X {
	//  return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg2.Start.Y}
	//}
	//
	//if seg2.Start.X > seg1.Start.X && seg2.Start.X < seg1.Stop.X {
	//  return true, image.Point{X: seg1.Stop.X - seg2.Start.X, Y: seg1.Start.Y}
	//}
	//
	//if seg1.Start.Y > seg2.Start.Y && seg1.Start.Y < seg2.Stop.Y {
	//  return true, image.Point{Y: seg2.Stop.Y - seg1.Start.Y, X: seg2.Start.X}
	//}
	//if seg2.Start.Y > seg1.Start.Y && seg2.Start.Y < seg1.Stop.Y {
	//  return true, image.Point{Y: seg1.Stop.Y - seg2.Start.Y, X: seg1.Start.X}
	//}
	return false, image.Point{}
}

func opositeDirections(seg1 Segment, seg2 Segment) bool {
	return seg1.Polarisation != seg2.Polarisation
}

type Wire struct {
	Segments []Segment
}

func Distance(w1 Wire, w2 Wire) int {
	distances := make([]int, 0, 10)

	for _, seg1 := range w1.Segments {
		for _, seg2 := range w2.Segments {
			intersects, point := seg1.intersects(seg2)
			if intersects {
				distances = append(distances, int(math.Abs(float64(point.X))+math.Abs(float64(point.Y))))
			}
		}
	}
	sort.Ints(distances)
	if len(distances) == 0 {
		return -1
	}
	return distances[0]
}

func NewWire(wire string) Wire {
	newWire := Wire{Segments: make([]Segment, 0, 10)}

	start := image.Point{}

	for i, s := range strings.Split(wire, ",") {
		newWire.Segments = append(newWire.Segments, NewSegment(s, start))
		start = newWire.Segments[i].Stop
	}
	return newWire
}

//
//func (w Wire) Distance() (x int, y int) {
// for _, segment := range w.Segments {
//   if segment.Direction == 'R' {
//     x = x + segment.Length
//   } else if segment.Direction == 'L' {
//     x = x - segment.Length
//   } else if segment.Direction == 'U' {
//     y = y + segment.Length
//   } else if segment.Direction == 'D' {
//     y = y - segment.Length
//   }
// }
// return
//}
