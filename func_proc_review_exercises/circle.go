package main

import "fmt"

func main() {
	var x, y int
	var c1x, c1y, c1r int
	var c2x, c2y, c2r int
	var distanceVal1, distanceVal2 int

	fmt.Scan(&c1x, &c1y, &c1r)
	fmt.Scan(&c2x, &c2y, &c2r)
	fmt.Scan(&x, &y)

	distanceVal1 = distance(x, y, c1x, c1y)
	distanceVal2 = distance(x, y, c2x, c2y)

	if iSinside(distanceVal1, c1r) && iSinside(distanceVal2, c2r) {
		fmt.Print("point is inside circle 1 and 2")
	} else if iSinside(distanceVal1, c1r) {
		fmt.Print("point is inside circle 1")
	} else if iSinside(distanceVal2, c2r) {
		fmt.Print("point is inside circle 2")
	} else {
		fmt.Print("point is outside circle 1 and 2")
	}

}

func distance(x, y, cx, cy int) int {
	return (x-cx)*(x-cx) + (y-cy)*(y-cy) //formula

}
func iSinside(distanceVal, r int) bool {
	return distanceVal <= r*r //condition

}
