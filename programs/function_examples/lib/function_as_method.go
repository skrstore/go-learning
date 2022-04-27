package lib

import (
	"fmt"
	"math"
)

func RunFunctionAsMethod() {
	circle := Circle{x: 0, y: 0, radius: 3}
	fmt.Printf("Circle Area : %f\n", circle.area())
}

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
