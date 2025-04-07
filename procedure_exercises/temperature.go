package main

import "fmt"

func main() {
	var reamur, fahrenheit, kelvin, celcius float64
	fmt.Scan(&celcius)
	tempConversion(celcius, &reamur, &fahrenheit, &kelvin)
	fmt.Printf("%.0fR %.0fF %.2fK", reamur, fahrenheit, kelvin)

}

func tempConversion(celcius float64, reamur, fahrenheit, kelvin *float64) {
	*reamur = celcius * 4 / 5
	*fahrenheit = celcius*9/5 + 32
	*kelvin = celcius + 273.15
}
