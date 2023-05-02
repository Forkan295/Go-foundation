package main

import "fmt"

type shape interface {
	getArea() float64
}

type circle struct {
	radius int
}

type square struct {
	width  int
	height int
}

func (c circle) getArea() float64 {
	return 3.14 * float64(c.radius)
}

func (s square) getArea() float64 {
	return float64(s.width * s.height)
}

func calculate(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{
		width:  10,
		height: 20,
	}

	cir := circle{
		radius: 323,
	}

	calculate(sq)
	calculate(cir)
}
