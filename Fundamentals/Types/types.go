package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 20, 100) // INT

	// only positive numbers uint8 (byte), uint16 (short), uint32 (int), uint64 (long)
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	// with sings int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("The maximum int64 value is", i1, "and the type is", reflect.TypeOf(i1))

	// mapping an unicode table (int32)
	var i2 rune = 'a'
	fmt.Println(i2, reflect.TypeOf(i2))

	// real numbers (float32, float64)
	var x float32 = 49.99
	fmt.Println(reflect.TypeOf(x))     // you can declare the type to be 32, buuuuuut
	fmt.Println(reflect.TypeOf(49.99)) // if you dont declare it, it will be float64

	// boolean type is bool
	fmt.Println(reflect.TypeOf(true))

	// Strings has the type string and char not exist, it will return int32
	var str string = "Hello, World!"
	fmt.Println("The type is", reflect.TypeOf(str), "and the length is", len(str))
	//String with multiple lines
	var str2 string = `Hello
	,
	World
	!`
	fmt.Println("The type is", reflect.TypeOf(str2), "and the length is", len(str2))
	fmt.Println(str2)
}
