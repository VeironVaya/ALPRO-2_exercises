package main

import "fmt"

func main() {
	var w, h, scale int
	var symbol byte

	fmt.Scanln(&w, &h)
	fmt.Scanf("%c %d", &symbol, &scale)

	if symbol == '+' {
		zoomIn(w, scale) 
		zoomIn(h, scale)
	} else if symbol == '-' {
		zoomOut(w, scale)
		zoomOut(h, scale)

	}

}

func zoomIn(num, scale int)  {
	fmt.Print( num * scale)
}

func zoomOut(num, scale int)  {
	fmt.Print( num / scale)
}
