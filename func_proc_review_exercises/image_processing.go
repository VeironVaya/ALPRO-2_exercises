package main

import "fmt"

func main() {
	var w, h, scale int
	var symbol byte

	fmt.Scanln(&w, &h)
	fmt.Scanf("%c %d", &symbol, &scale)

	if symbol == '+' {
		fmt.Print(zoomIn(w, scale), zoomIn(h, scale))
	} else if symbol == '-' {
		fmt.Print(zoomOut(w, scale), zoomOut(h, scale))

	}

}

func zoomIn(num, scale int) int {
	return num * scale
}

func zoomOut(num, scale int) int {
	return num / scale
}
