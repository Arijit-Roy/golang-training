package main

import (
	"fmt"
)

type shape interface {
	getArea() (string, float64)
}

type square struct {
	sideLength float64
	stype      string
}

type triangle struct {
	base   float64
	height float64
	stype  string
}

func main() {
	sq := square{
		sideLength: 15,
		stype:      "square",
	}

	tra := triangle{
		base:   5.5,
		height: 2.6,
		stype:  "triangle",
	}

	printArea(sq)
	printArea(tra)
}

func (t triangle) getArea() (string, float64) {
	return t.stype, 0.5 * t.base * t.height
}

func (s square) getArea() (string, float64) {
	return s.stype, s.sideLength * s.sideLength
}

func printArea(s shape) {
	stype, area := s.getArea()

	fmt.Println(stype, "has area of", area)
}
