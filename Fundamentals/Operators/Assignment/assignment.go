package main

import "fmt"

func main(){
	// Using that way you need to explicitly show the type
	var b byte = 3
	fmt.Println(b)

	// And that way it takes the type of the value that are being assigned and set to that variable
	i := 3
	fmt.Println(i)

	i += 3 // i = i + 3
	i -= 3 //  i = i - 3
	i /= 3 // i =  i / 3
	i *= 3 //  i =  i * 3
	i %= 3 // i = i % 3
	fmt.Println(i)

	x, y := 1,2
	fmt.Println(x,y)

	x, y = y, x
	fmt.Println(x,y)

}