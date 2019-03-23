package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum => ", a+b)
	fmt.Println("Subtraction => ", a-b)
	fmt.Println("Division => ", a/b)
	fmt.Println("Multiplication => ", a*b)
	fmt.Println("Module => ", a%b)


	// Bitwise
	fmt.Println("And => ", a&b)  // 11 & 10 = 10
	fmt.Println("Or => ", a|b)  // 11 & 10 = 11
	fmt.Println("XOR => ", a^b)  // 11 & 10 = 01


	c := float64(a)
	d := float64(b)

	// Another operations using math, and normally needs to be float64 instead of int
	fmt.Println("Biggest => ", math.Max(c,d))
	fmt.Println("Lowest => ", math.Max(c,d))
	fmt.Println("Exponential => ", math.Pow(c,d))
}
