package main

import (
	"fmt"
)

type Geometry struct {
	area      int
	perimeter int
}

type Rectangle struct {
	length   int
	width    int
	color    string
	property Geometry
}

func inputData(rect *Rectangle) {
	fmt.Scan(&rect.length)

	fmt.Scan(&rect.width)

	fmt.Scan(&rect.color)
}

func calculate(rect *Rectangle) {
	rect.property.area = rect.length * rect.width
	rect.property.perimeter = 2 * (rect.length + rect.width)
}

func main() {
	var rect Rectangle

	inputData(&rect)
	calculate(&rect)

	fmt.Println("Length   :", rect.length)
	fmt.Println("Width    :", rect.width)
	fmt.Println("Color    :", rect.color)
	fmt.Println("Area     :", rect.property.area)
	fmt.Println("Perimeter:", rect.property.perimeter)
}
