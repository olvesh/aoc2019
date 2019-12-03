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

func (s Segment) ManhattanDistance(point image.Point) int {
	//|x1 – x2| + |y1 – y2|
	return int(math.Abs(float64((s.Stop.X - point.X) + (s.Stop.Y - point.Y))))
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

func Steps(w1 Wire, w2 Wire) int {
	steps := make([]int, 0, 10)

	w1Steps := 0
	for i, seg1 := range w1.Segments {
		w1Steps = w1Steps + seg1.Length

		w2Steps := 0
		for j, seg2 := range w2.Segments {

			w2Steps = w2Steps + seg2.Length
			intersects, p := seg1.intersects(seg2)
			if intersects {
				//|x1 – x2| + |y1 – y2|

				distance1 := seg1.ManhattanDistance(p)
				distance2 := seg2.ManhattanDistance(p)
				numSteps := w1Steps + w2Steps - distance1 - distance2
				log.Printf("Found intersection at %v", p)
				log.Printf("Found intersection idx  i:%v j:%v", i, j)
				steps = append(steps, numSteps)
			}
		}
	}
	sort.Ints(steps)
	if len(steps) == 0 {
		return -1
	}
	return steps[0]
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
