package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{height: 5, base: 5}
	sq := square{sideLength: 55}

	printArea(tr)
	printArea(sq)
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height / 2
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
