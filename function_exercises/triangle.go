/*Triangle Inequality Theorem, which states that the sum of two
side lengths of a triangle is always greater than the third side.*/

package main

import "fmt"

func main() {
	var x, y, z int
	var isTriangleVal bool
	
	fmt.Scan(&x, &y, &z)
	isTriangleVal = isTriangle(x, y, z)
	if isTriangleVal {
		fmt.Print("triangle")
	} else {
		fmt.Print("not triangle")
	}

}

func isTriangle(x, y, z int) bool {
	return x+y > z && x+z > y && z+y > x

}
