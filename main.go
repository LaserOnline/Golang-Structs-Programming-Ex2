package main

import (
	"fmt"
	"math"
)

type dataArea struct {
	width  float64
	height float64
	radius float64
}

func main() {
	d := dataArea{
		width:  4.0,
		height: 2.5,
		radius: 2,
	}

	x := squareArea(d)
	z := radiusArea(&d)
	fmt.Println("Value ->", x)
	fmt.Println("Value ->", z)
}

func squareArea(m dataArea) float64 {
	result := m.width * m.height
	return result
}

func radiusArea(m *dataArea) float64 {
	result := math.Pi * m.radius * m.radius
	return result
}
