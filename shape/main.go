package main

import (
	"fmt"
	"strings"
)

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

	//triangle := triangle{height: 10, base: 3}
	triangle := triangle{}
	square := square{}

	fmt.Println("Enter the shape name: Square or Triangle")
	fmt.Println("---------------------")

	var text string
	fmt.Scanf("%s\n", &text)

	if strings.Compare("Triangle", text) == 0 {
		printArea(triangle)
	} else {
		printArea(square)
	}

}

func printArea(s shape) {
	fmt.Printf("Area %f", s.getArea())
}

func (square) getArea() float64 {

	var l float64
	fmt.Println("Enter sideLength of Square")

	_, err := fmt.Scanf("%f\n", &l)
	if err != nil {
		fmt.Println(err)
	}
	return l * l
}

func (triangle) getArea() float64 {
	var triangleBase float64
	var triangleHeight float64
	var areaOfTriangle float64
	fmt.Println("Enter the base of triangle:")
	fmt.Scanf("%f\n", &triangleBase)
	fmt.Println("Enter the height of triangle:")
	fmt.Scanf("%f\n", &triangleHeight)
	areaOfTriangle = (triangleBase * triangleHeight) / 2
	return areaOfTriangle
}
