package main

import "fmt"

func main() {
	var hours, minutes, seconds int
	var conversionVal int

	fmt.Scan(&hours, &minutes, &seconds)
	conversionVal = timeConversion(hours, minutes, seconds)
	fmt.Printf("Converted value = %d seconds", conversionVal)

}

func timeConversion(h, m, s int) int {
	return h*3600 + m*60 + s

}
