package main

import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (react Rectangle) area() float32 {
	return react.length * react.breadth
}

func getArea(shape Shape) float32 {
	return shape.area()
}

func main() {
	reactangle := Rectangle{2.5, 5.0}
	fmt.Printf("Area of rectangle l=%f b=%f = %f\n", reactangle.length, reactangle.breadth, getArea(reactangle))
}
