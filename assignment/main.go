package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	baseLength float64
}

func main() {
	tr := triangle{3, 4}
	sqr := square{4}

	printArea(tr)
	printArea(sqr)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.height * tr.base
}

func (sqr square) getArea() float64 {
	return sqr.baseLength * sqr.baseLength
}
