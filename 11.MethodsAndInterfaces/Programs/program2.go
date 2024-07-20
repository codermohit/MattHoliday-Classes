package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

/* Path slice contains the points followed by a line. We want the distance covered by following those points */
type Path []Point

/* Method to calculate distance between two points */
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

/* Method to scale the line, either use a pointer receiver or  */
func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

/* Distance method for the Path slice, either use Pointer or return the value in case of value receiver */
func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return
}

func main() {
	side := Line{Point{1, 2}, Point{4, 5}}
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(side.ScaleBy(2).Distance())
	fmt.Printf("Perimeter : %v\n", perimeter.Distance())
}
