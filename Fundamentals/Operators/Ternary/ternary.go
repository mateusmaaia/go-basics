package main

import "fmt"

// There isnt the ternary operator :(
// But we will show some options

func getResult(grade float64) string {
	if grade >= 6 {
		return "Approved"
	}

	return "Disapproved"

	// return grade >= 6 ? "Approved" : "Disapproved"
	// :(
}


func main () {
	fmt.Println(getResult(6.2))
}
