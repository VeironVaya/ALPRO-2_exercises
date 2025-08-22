package main

import "fmt"

const pi float64 = 3.14

func main() {
	var r, h1, h2, d1, d2 float64
	var circleAreaVal float64
	var volumeVal1, massVal1 float64
	var volumeVal2, massVal2 float64

	fmt.Scan(&r)
	fmt.Scan(&h1, &d1)
	fmt.Scan(&h2, &d2)
	
	circleAreaVal = circleArea(r)
	volumeVal1 = volume(circleAreaVal, h1)
	volumeVal2 = volume(circleAreaVal, h2)
	massVal1 = mass(volumeVal1, d1)
	massVal2 = mass(volumeVal2, d2)
	display(massVal1, massVal2)

}

func circleArea(r float64) float64 {
	return r * r * pi

}

func volume(circleAreaVal, h float64) float64 {
	return circleAreaVal * h

}

func mass(volumeVal, density float64) float64 {
	return volumeVal * density
}

func display(massVal1, massVal2 float64) {

	var difference float64

	difference = massVal1 - massVal2
	if difference < 0 {
		difference = difference * -1
	}

	if difference == 0 {
		fmt.Print("BALANCE")
	} else {
		fmt.Printf("%.3f", difference)
	}

}
