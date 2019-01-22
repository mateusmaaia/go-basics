package main

import "fmt"

func main() {
	fmt.Print("Same ")
	fmt.Print("Line!\n\n")

	fmt.Println("New")
	fmt.Println("Line!")

	PI := 3.1415
	fmt.Println("The value of PI is " + fmt.Sprint(PI))
	fmt.Println("The value of PI is ", PI)
	fmt.Printf("The value of PI is %.6f", PI)

	a := 1
	b := 1.9
	c := true
	d := "wow"
	fmt.Printf("\n%d %f %4.f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
