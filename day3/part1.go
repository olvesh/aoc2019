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
	intersects, point := segmentsIntersects(s, seg2)
	if (point == image.Point{}) {
		return false, image.Point{}
	}
	return intersects, point

	////if opositeDirections(s, seg2) {
	////
	//if b, point := segmentsIntersects(s, seg2); b {
	//	return b, point
	//}
	//if b, point := segmentsIntersects(seg2, s); b {
	//	return b, point
	//}
	////}
	//return false, image.Point{}
}

func segmentsIntersects(seg1 Segment, seg2 Segment) (bool, image.Point) {
	if seg1.Polarisation == hor && seg2.Polarisation == ver {
		minY, maxY := seg2.minMaxY()
		if seg1.Start.Y >= minY && seg1.Start.Y <= maxY {
			minX, maxX := seg1.minMaxX()
			if minX <= seg2.Start.X && maxX >= seg2.Start.X {
				return true, image.Point{seg2.Start.X, seg1.Start.Y}
			}
		}
	}

	if seg1.Polarisation == ver && seg2.Polarisation == hor {
		minX, maxX := seg2.minMaxX()
		if seg1.Start.X >= minX && seg1.Start.X <= maxX {
			minY, maxY := seg1.minMaxY()
			if minY <= seg2.Start.Y && maxY >= seg2.Start.Y {
				return true, image.Point{seg1.Start.X, seg2.Start.Y}
			}
		}
	}
	return false, image.Point{}
}

func (s Segment) minMaxX() (min int, max int) {
	min = int(math.Min(float64(s.Start.X), float64(s.Stop.X)))
	max = int(math.Max(float64(s.Start.X), float64(s.Stop.X)))
	return
}
func (s Segment) minMaxY() (min int, max int) {
	min = int(math.Min(float64(s.Start.Y), float64(s.Stop.Y)))
	max = int(math.Max(float64(s.Start.Y), float64(s.Stop.Y)))
	return
}

//func segmentsIntersect(seg1 Segment, seg2 Segment) (bool, image.Point) {
//	// x axis muyst be within y
//
//
//
//	if seg1.Start.X <= seg2.Stop.X && seg2.Start.X <= seg1.Stop.X &&
//		seg1.Start.Y >= seg2.Stop.Y {
//
//		return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg1.Start.Y}
//	}
//	if seg1.Start.Y <= seg2.Stop.Y && seg2.Start.Y <= seg1.Stop.Y {
//		return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg1.Start.Y}
//	}
//
//	//if seg1.Start.Y < seg2.Stop.Y && seg2.Start.Y < seg1.Stop.Y {
//	//	return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg2.Start.Y}
//	//}
//
//	//if seg1.Start.X > seg2.Start.X && seg1.Start.X < seg2.Stop.X {
//	//  return true, image.Point{X: seg2.Stop.X - seg1.Start.X, Y: seg2.Start.Y}
//	//}
//	//
//	//if seg2.Start.X > seg1.Start.X && seg2.Start.X < seg1.Stop.X {
//	//  return true, image.Point{X: seg1.Stop.X - seg2.Start.X, Y: seg1.Start.Y}
//	//}
//	//
//	//if seg1.Start.Y > seg2.Start.Y && seg1.Start.Y < seg2.Stop.Y {
//	//  return true, image.Point{Y: seg2.Stop.Y - seg1.Start.Y, X: seg2.Start.X}
//	//}
//	//if seg2.Start.Y > seg1.Start.Y && seg2.Start.Y < seg1.Stop.Y {
//	//  return true, image.Point{Y: seg1.Stop.Y - seg2.Start.Y, X: seg1.Start.X}
//	//}
//	return false, image.Point{}
//}

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
