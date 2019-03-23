package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Only postfix
	x++ // x += 1 || x + 1
	fmt.Println(x)

	y-- // y -= 1 || y - 1
	fmt.Println(y)

	// Not exist fmt.Println(x == y--)
}
