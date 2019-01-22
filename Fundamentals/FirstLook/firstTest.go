// Executable programs always starts by main package
package main

// Use import to use external packages, you can list it as an array
import (
	"fmt"
)

// Executable program always use function main, and it can only be declared once
func main() {
	var text string = "world"

	fmt.Printf("Hello, %s!\n", text)
}
