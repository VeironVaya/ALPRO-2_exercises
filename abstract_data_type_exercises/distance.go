package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {
	var p1, p2 point
	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)
	fmt.Print(distance(p1, p2))

}

func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}
