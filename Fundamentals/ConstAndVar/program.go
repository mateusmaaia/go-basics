package main

import (
	"fmt"
	m "math" // You can put a word before the package name to use as Alias and you can use an "_" before the package to use it indirectly
	"reflect"
)

func main() {
	// If you are declaring the value, it's not obligatory to declare types, it will get the value type
	const PI float64 = 3.1415
	var radius = 3.2 // Type float64 inferred

	// Reduced way to declare var
	area := PI * m.Pow(radius, 2)
	fmt.Println("Area:", area)

	// Another way to declare const and var

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "wow!"
	fmt.Println(g, h, i)

	fmt.Println(reflect.TypeOf(area), reflect.TypeOf(PI), reflect.TypeOf(radius), reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d), reflect.TypeOf(e), reflect.TypeOf(f), reflect.TypeOf(g), reflect.TypeOf(h), reflect.TypeOf(i))
}
