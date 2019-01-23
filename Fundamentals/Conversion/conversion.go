package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // You need to use same numerical types

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// Be careful, if you are converting int to string it will use the relative ASCII character
	fmt.Println("Test " + string(123))

	// Int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// String to int, that function return 2 values, the first will be a number and the second one will be an error (you can use _ if you wont deal with it)
	num, er := strconv.Atoi("123")
	println(num, er)

	// String to boolean
	boolean, er := strconv.ParseBool("123true")
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println(boolean)
	}
}
